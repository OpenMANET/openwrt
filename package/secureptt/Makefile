include $(TOPDIR)/rules.mk

PKG_NAME:=pttcomm
PKG_VERSION:=git-7adb61c5
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://codeberg.org/48554e6d/pttcomm.git
PKG_SOURCE_VERSION:=7adb61c5b95aedf2fd97536d44459c7c011d909c
PKG_MIRROR_HASH:=skip

PKG_LICENSE:=GPL-2.0-only
PKG_MAINTAINER:=info@openmanet.net
PKG_BUILD_PARALLEL:=1

include $(INCLUDE_DIR)/package.mk

define Package/pttcomm
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=PTT audio communication tool (GStreamer-based)
  DEPENDS:=+glib2 +gstreamer1-libs +gstreamer1-plugins-base +gstreamer1
endef

define Package/pttcomm/description
pttcomm: a simple push-to-talk audio utility that uses GStreamer.
Built from https://codeberg.org/48554e6d/pttcomm.git (commit 7adb61c5).
endef

# Let OpenWrt's wrappers provide pkg-config, CFLAGS, LDFLAGS, etc.
define Build/Compile
	$(MAKE) -C $(PKG_BUILD_DIR) \
		CC="$(TARGET_CC)" \
		CFLAGS="$(TARGET_CFLAGS) $(TARGET_CPPFLAGS)" \
		LDFLAGS="$(TARGET_LDFLAGS)"
endef

define Package/pttcomm/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/pttcomm $(1)/usr/bin/
endef

$(eval $(call BuildPackage,pttcomm))