#!/bin/bash
# -----------------------------------------------------------------------------
# morse-build: Build OpenWrt and sign sysupgrade images
#
# Modes:
#   localsigned Build + sign using a locally generated keypair.
#   signed      Build + sign using a user-provided keypair (requires --key/--pubkey).
#   sign-only   Append a signature trailer to an already-built image
#               (requires --image, --key, --pubkey).
#
# Signing vs verification:
# - Signing attaches a usign/ucert/fwtool trailer using your *private* key.
# - Verification happens later, on-device during sysupgrade, using the *public* key
#   baked into the running image via morse-firmware-sign package.
#   Using sign-only with a different private key requires that the matching public
#   key is present in the *running* image for verification to succeed.
# -----------------------------------------------------------------------------

# ---------- Defaults ----------
MORSE_FW_SIGNING_PKG="CONFIG_PACKAGE_morse-firmware-sign"
CONFIG_SIG_CHECK="CONFIG_SIGNATURE_CHECK"

# Host tools path (used for --sign local path)
STAGING_BIN_DEFAULT="./staging_dir/host/bin"
STAGING_BIN="${STAGING_BIN_DEFAULT}"
export PATH="$STAGING_BIN_DEFAULT:$PATH"

MODE="build-signed-local" # default if no mode provided
IMG_NAME=""
PRIV_KEY=""           # --key (private)
PUB_KEY=""            # --pubkey (public)
JOBS=6
VERBOSE=0

SIGNED_SUFFIX="${SIGNED_SUFFIX:-}"   # only used for local --sign flow
CLEANUP_KEYS=0

# ---------- Helpers ----------
die()  { echo "ERROR: $*" >&2; exit 1; }
note() { echo "[morse-build] $*"; }

usage() {
cat <<'EOF'
morse-build - Build OpenWrt and sign sysupgrade images

Usage:
  ./morse-build {localsigned | signed | sign-only} [options]
                 [--image NAME] [--key PRIV] [--pubkey PUB]
                 [--jobs N] [--verbose]

Modes:
  localsigned   Build and sign using a locally generated keypair (created while building).
  signed        Build and sign using a provided keypair.
                Requires: --key PRIV and --pubkey PUB
  sign-only     Sign ONE existing image (in-place) using a provided keypair.
                Requires: --image NAME (under bin/targets/...), --key PRIV, --pubkey PUB

Notes:
- To have the *device* enforce signature checks during sysupgrade, include:
    morse-firmware-sign package. 
  (By default this installs the Morse release public key. To verify against a
   custom key, include that public key in the image via overlay or a customized
   package.)
EOF
}

in_top_of_openwrt(){ [[ -f "Makefile" && -f "include/target.mk" && -x "scripts/feeds" ]]; }
ensure_in_top(){ in_top_of_openwrt || die "Run from the OpenWrt top directory (with Makefile and scripts/feeds)."; }
cfg_str(){ sed -nr "s/^$1=\"([^\"]*)\"/\\1/p" .config; }
cfg_check(){ grep -qE "^$1=y" .config 2>/dev/null; }

detect_target_bin_dir(){
	[[ -f ".config" ]] || die "No .config found. Run 'make menuconfig' first."
	local board sub
	board="$(cfg_str CONFIG_TARGET_BOARD)"
	sub="$(cfg_str CONFIG_TARGET_SUBTARGET)"
	[[ -n "$board" && -n "$sub" ]] || die "Could not read CONFIG_TARGET_BOARD/CONFIG_TARGET_SUBTARGET from .config"
	echo "bin/targets/${board}/${sub}"
}

resolve_tool() {
	local name="$1"
	local cand="${STAGING_BIN}/${name}"

	if [[ -x "$cand" ]]; then
		echo "$cand"
		return 0
	fi

	if command -v "$name" >/dev/null 2>&1; then
		command -v "$name"
		return 0
	fi

	die "Tool not found: $name (searched in $STAGING_BIN and PATH)"
}

stage_keys() {
	# If user supplied the keys, stage both of them for signing.
	install -m 644 "$PRIV_KEY" ./key-build       || die "Failed to stage private key to ./key-build"
	install -m 644 "$PUB_KEY"  ./key-build.pub   || die "Failed to stage public key to ./key-build.pub"
	CLEANUP_KEYS=1
	note "Using user-provided keys for signing (staged to ./key-build and ./key-build.pub)."
}

sign_image() {
	# Sign one image in place using usign/ucert/fwtool.
	local img="$1" key="$2" pub="$3"

	local USIGN UCERT FWTOOL
	USIGN="$(resolve_tool usign)"
	UCERT="$(resolve_tool ucert)"
	FWTOOL="$(resolve_tool fwtool)"

	note "Signing (local) image: $(basename "$img")"
	note "  usign:  $USIGN"
	note "  ucert:  $UCERT"
	note "  fwtool: $FWTOOL"

	# Strip any existing trailer
	local tmp sig cert extracted stripped
	tmp="$(mktemp)"
	if "$FWTOOL" -q -s "$tmp" "$img" && [[ -s "$tmp" ]]; then
		note "  stripping existing metadata from image"
		"$FWTOOL" -q -t -s "$tmp" "$img"
	fi
	rm -f "$tmp"

	# Sign the unsigned content and wrap into ucert
	sig="${img}.sig"
	cert="${img}.ucert"
	: >"$cert"
	"$USIGN" -S -m "$img" -s "$key" -x "$sig"
	"$UCERT" -A -c "$cert" -x "$sig"

	# Append trailer to the image
	"$FWTOOL" -S "$cert" "$img"

	# Optional verify (if a public key file is given)
	extracted="${img}.extracted.ucert"
	"$FWTOOL" -s "$extracted" "$img" >/dev/null
	if [[ -s "$extracted" && -n "$pub" && -f "$pub" ]]; then
		stripped="$(mktemp)"
		note "  verify step 1/2: extract message"
		if "$FWTOOL" -T -s /dev/null "$img" > "$stripped"; then
			note "  verify step 2/2: ucert verify using pubkey"
			if "$UCERT" -V -m "$stripped" -c "$extracted" -p "$pub" >/dev/null 2>&1; then
				note "  verify: ok"
			else
				note "  verify: FAILED (ucert -V)"
			fi
		else
			note "  verify: FAILED (fwtool -T)"
		fi
		rm -f "$stripped"
	else
		note "  (info) local verify skipped (no public key provided or empty extract)"
	fi

	rm -f "$sig" "$cert" "$extracted"
	note "✔ done: $(basename "$img")"
}

cleanup_staged_keys() {
	[[ "${CLEANUP_KEYS:-1}" -eq 1 ]] || return 0

	for key in ./key-build ./key-build.pub; do
		[[ -f "$key" ]] && rm -f "$key"
	done
}

# ---------- Arg parse ----------
# 1) Mode (positional)
if [[ $# -gt 0 ]]; then
	case "$1" in
		localsigned) MODE="build-signed-local"; shift ;;
		signed)      MODE="build-signed";    shift ;;
		sign-only)   MODE="sign";            shift ;;
		-h|--help) usage; exit 0 ;;
		--*)       : ;;  # no explicit mode → default to unsigned; continue parsing options
		*)         die "Unknown mode: $1 (use: localsigned | signed | sign-only, or --help)";;
	esac
fi

# Options
while [[ $# -gt 0 ]]; do
	case "$1" in
		--image)   IMG_NAME="${2:-}"; [[ -n "$IMG_NAME" ]] || die "--image requires a filename"; shift 2 ;;
		--key)     PRIV_KEY="${2:-}"; [[ -n "$PRIV_KEY" ]] || die "--key requires a file path"; shift 2 ;;
		--pubkey)  PUB_KEY="${2:-}";  [[ -n "$PUB_KEY"  ]] || die "--pubkey requires a file path"; shift 2 ;;
		--jobs)    JOBS="${2:-}"; [[ "$JOBS" =~ ^[0-9]+$ ]] || die "--jobs requires a number"; shift 2 ;;
		--verbose) VERBOSE=1; shift ;;
		-h|--help) usage; exit 0 ;;
		*)         die "Unknown option: $1 (see --help)";;
	esac
done

# ---------- Ensure that the script is run from the top openwrt directory ----------
ensure_in_top
[[ -f ".config" ]] || die "No .config found. Run 'make menuconfig' first."

if [[ "$MODE" == "build-signed" || "$MODE" == "sign" ]]; then
	[[ -n "$PRIV_KEY" && -n "$PUB_KEY" ]] || die "signed mode requires --key and --pubkey"
	cfg_check "$MORSE_FW_SIGNING_PKG" || die "Signing package not enabled: ${MORSE_FW_SIGNING_PKG}=y not found in .config"
	cfg_check "$CONFIG_SIG_CHECK"     || die "Signature check not enabled: ${CONFIG_SIG_CHECK}=y not found in .config"
	stage_keys
fi

# ---------- Execute ----------
if [[ "$MODE" == "sign" ]]; then
	[[ -n $IMG_NAME ]] || die "Image name must be specified using --image option."
	TARGET_BIN_DIR="$(detect_target_bin_dir)"
	[[ -f "$TARGET_BIN_DIR/$IMG_NAME" ]] || die "Image: $IMG_NAME not available under $TARGET_BIN_DIR"
	note "Mode: sign (local). Target: $TARGET_BIN_DIR/$IMG_NAME"
	sign_image "$TARGET_BIN_DIR/$IMG_NAME" "$PRIV_KEY" "$PUB_KEY"
	note "Done (sign)."
	exit 0
fi

# Build signed firmware via native OpenWrt signing method
if [[ "$MODE" == "build-signed" ]]; then
	note "Build OpenWrt Firmware and do native signing with given keys."
else
	note "Build OpenWrt Firmware and do native signing with locally generated keys."
fi

note "Starting build… (logs at logs/build.log; view with: tail -f ./logs/build.log)"
mkdir -p ./logs
MAKE_ARGS=(-j"${JOBS:-6}")
if [[ ${VERBOSE:-0} -eq 1 ]]; then
	MAKE_ARGS+=("V=sc")
else
	MAKE_ARGS+=("-s")
fi

make "${MAKE_ARGS[@]}" 2>&1 | tee logs/build.log

cleanup_staged_keys

note "Build complete."