include $(TOPDIR)/rules.mk

PKG_NAME:=openmanetd
PKG_VERSION:=0.0.4
PKG_RELEASE:=1

PKG_SOURCE:=$(PKG_NAME)-$(PKG_VERSION).tar.gz
PKG_SOURCE_URL:=https://codeload.github.com/openmanet/openmanetd/tar.gz/$(PKG_VERSION)?
PKG_SOURCE_VERSION:=$(PKG_VERSION)
PKG_HASH:=skip

PKG_LICENSE:=MIT
PKG_MAINTAINER:=Corey Wagehoft

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_BUILD_FLAGS:=no-mips16

GO_PKG:=github.com/openmanet/openmanetd
GO_PKG_LDFLAGS_X:=cmd.Version=$(PKG_VERSION)

include $(INCLUDE_DIR)/package.mk
include ../../../feeds/packages/lang/golang/golang-package.mk

# Explicit runtime deps (correct OpenWrt syntax)
RDEPENDS:$(PKG_NAME) += libopusfile

define Package/openmanetd
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=OpenMANET Manager
  DEPENDS:=$(GO_ARCH_DEPENDS) +kmod-batman-adv +alfred +portaudio +libopus +libevdev +alsa-lib +kmod-usb-core +kmod-usb-audio +kmod-sound-core +libopusfile
endef

define Package/openmanetd/description
A management application for openmanetd written in Go.
Includes:
- pttgo: A multicast push-to-talk voice application written in Go using PortAudio, Opus, and evdev.
endef

define Package/openmanetd/conffiles
/etc/openmanetd/
endef

define Package/openmanetd/install
	$(call GoPackage/Package/Install/Bin,$(1))

	$(INSTALL_DIR) $(1)/etc/openmanetd/
	$(INSTALL_CONF) $(CURDIR)/files/sample_config.yml $(1)/etc/openmanetd/config.yml
	$(INSTALL_DIR) $(1)/etc/config/
	$(INSTALL_CONF) $(CURDIR)/files/openmanetd.config $(1)/etc/config/openmanetd
	$(INSTALL_DIR) $(1)/etc/init.d/
	$(INSTALL_BIN) $(CURDIR)/files/openmanetd.init $(1)/etc/init.d/openmanetd
endef

$(eval $(call GoBinPackage,openmanetd))
$(eval $(call BuildPackage,openmanetd))
