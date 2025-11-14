include $(TOPDIR)/rules.mk
include $(INCLUDE_DIR)/package.mk

PKG_NAME:=ptt-go
PKG_VERSION:=1.0.0
PKG_RELEASE:=1
PKG_LICENSE:=MIT
PKG_MAINTAINER:=You <you@example.com>

PKG_SOURCE_PROTO:=none
PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)-$(PKG_VERSION)
PKG_INSTALL:=1

define Package/ptt-go
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=Multicast PTT voice app
  DEPENDS:=+portaudio +libopus +libevdev +alsa-lib +kmod-usb-core +kmod-usb-audio +kmod-sound-core +libopusfile
endef

# Explicit runtime deps (correct OpenWrt syntax)
RDEPENDS:$(PKG_NAME) += libopusfile


define Package/ptt-go/description
A multicast push-to-talk voice application written in Go using PortAudio, Opus, and evdev.
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
	$(CP) ./src/* $(PKG_BUILD_DIR)/
endef

define Build/Compile
	cd $(PKG_BUILD_DIR)/github.com/ptt/ptt-go && \
		GOOS=linux GOARCH=arm64 CGO_ENABLED=1 \
		CC=$(TARGET_CC) \
		PKG_CONFIG_PATH="$(STAGING_DIR)/usr/lib/pkgconfig:$(STAGING_DIR)/usr/share/pkgconfig" \
		CGO_CFLAGS='$(TARGET_CFLAGS) -I$(STAGING_DIR)/usr/include' \
		CGO_LDFLAGS='$(TARGET_LDFLAGS) -L$(STAGING_DIR)/usr/lib -Wl,-rpath-link=$(STAGING_DIR)/usr/lib -Wl,-rpath=/usr/lib' \
		go build -mod=vendor -v -trimpath -o $(PKG_BUILD_DIR)/ptt-go .
endef

define Package/ptt-go/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/ptt-go $(1)/usr/bin/ptt-go
endef

define Build/Configure
endef

define Build/InstallDev
endef

define Build/Install
endef

$(eval $(call BuildPackage,ptt-go))

