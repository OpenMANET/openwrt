include $(TOPDIR)/rules.mk

PKG_NAME:=adsbtocot
PKG_VERSION:=1
PKG_RELEASE:=2
PKG_MAINTAINER:=info@openmanet.net
PKG_LICENSE:=MIT
PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)
HOST_BUILD_DEPENDS:=python3/host
PKGARCH:=all

include $(INCLUDE_DIR)/package.mk

define Package/adsbtocot
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=ADS-B to Cursor-on-Target (ADSBCOT) for OpenWrt (dump1090 local)
  DEPENDS:=+python3 +dump1090
endef

define Package/adsbtocot/description
Installs ADSBCOT from PyPI and provides a procd service to forward ADS-B tracks
from dump1090 to CoT/TAK (multicast or unicast).
endef

# Persist user edits
define Package/adsbtocot/conffiles
/etc/adsbcot/adsbcot.ini
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
endef

# Vendor ADSBCOT as pure Python:
# - drop 'with_pymodes' to avoid numpy/ELF
# - disable aiohttp/yarl/multidict/frozenlist C-extensions
# - forbid wheels so pip builds pure-Python installs
define Build/Compile
	AIOHTTP_NO_EXTENSIONS=1 \
	YARL_NO_EXTENSIONS=1 \
	MULTIDICT_NO_EXTENSIONS=1 \
	FROZENLIST_NO_EXTENSIONS=1 \
	PIP_NO_BINARY=:all: \
	$(STAGING_DIR_HOSTPKG)/bin/pip3 install 'adsbcot' \
		--prefix=/usr --root='$(PKG_INSTALL_DIR)' --no-cache-dir --no-compile

	# Safety: ensure no native .so crept in
	find '$(PKG_INSTALL_DIR)/usr' -type f \( -name '*.so' -o -name '*.so.*' \) -print -delete || true
endef

define Package/adsbtocot/install
	# Config + init
	$(INSTALL_DIR) $(1)/etc/adsbcot $(1)/etc/init.d
	$(INSTALL_CONF) ./files/adsbcot.ini  $(1)/etc/adsbcot/adsbcot.ini
	$(INSTALL_BIN)  ./files/adsbcot.init $(1)/etc/init.d/adsbcot

	# Python site-packages and console scripts
	$(INSTALL_DIR) $(1)/usr
	cp -a $(PKG_INSTALL_DIR)/usr/* $(1)/usr/

	# Ensure no ELF in final image
	find '$(1)/usr' -type f \( -name '*.so' -o -name '*.so.*' \) -print -delete || true

	# Normalize console-script shebangs to the target Python
	[ -f "$(1)/usr/bin/adsbcot" ] && $(SED) '1s|^#!.*|#!/usr/bin/python3|' "$(1)/usr/bin/adsbcot" && chmod 0755 "$(1)/usr/bin/adsbcot"
	for f in pytak aircot; do \
		if [ -f "$(1)/usr/bin/$$f" ]; then \
			$(SED) '1s|^#!.*|#!/usr/bin/python3|' "$(1)/usr/bin/$$f"; \
			chmod 0755 "$(1)/usr/bin/$$f"; \
		fi; \
	done

	# Trim caches to save space
	find "$(1)/usr/lib/python3"* -name '__pycache__' -type d -exec rm -rf {} + 2>/dev/null || true
	find "$(1)/usr/lib/python3"* -type f -name '*.pyc' -delete 2>/dev/null || true

	# First-boot defaults (gain, ports, etc.)
	$(INSTALL_DIR) $(1)/etc/uci-defaults
	$(INSTALL_BIN) ./files/99-adsb-defaults $(1)/etc/uci-defaults/99-adsb-defaults
endef

$(eval $(call BuildPackage,adsbtocot))