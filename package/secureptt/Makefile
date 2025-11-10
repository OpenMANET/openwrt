include $(TOPDIR)/rules.mk

PKG_NAME:=pttcomm
PKG_RELEASE:=1

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://codeberg.org/48554e6d/pttcomm.git
PKG_SOURCE_VERSION:=7adb61c5b95aedf2fd97536d44459c7c011d909c
# You can pin a mirror hash after first successful fetch (preferred).
# Until then, skip to allow dev builds:
PKG_MIRROR_HASH:=skip

PKG_LICENSE:=GPL-2.0-only
PKG_MAINTAINER:=info@openmanet.net

PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)
PKG_BUILD_PARALLEL:=1

include $(INCLUDE_DIR)/package.mk

define Package/pttcomm
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=PTT audio communication tool (GStreamer-based)
  DEPENDS:=+gstreamer1-libs +gstreamer1-plugins-base
endef

define Package/pttcomm/description
pttcomm: a simple push-to-talk audio utility that uses GStreamer.
This package builds from the Codeberg repo and installs the 'pttcomm' binary.
endef

# If upstream Makefile honors CC/CFLAGS/LDLIBS via environment, this is enough.
# Add PKG_CONFIG, LDFLAGS, or extra INCLUDE paths if needed.
TARGET_CFLAGS += $(TARGET_CPPFLAGS)
TARGET_LDFLAGS +=

define Build/Prepare
	$(call Build/Prepare/Default)
endef

define Build/Compile
	$(MAKE) -C $(PKG_BUILD_DIR) \
		CC="$(TARGET_CC)" \
		CPPFLAGS="$(TARGET_CPPFLAGS)" \
		CFLAGS="$(TARGET_CFLAGS)" \
		LDFLAGS="$(TARGET_LDFLAGS)" \
		PKG_CONFIG="$(STAGING_DIR_HOST)/bin/pkg-config" \
		PKG_CONFIG_PATH="$(STAGING_DIR)/usr/lib/pkgconfig:$(STAGING_DIR)/usr/share/pkgconfig"
endef

define Package/pttcomm/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/pttcomm $(1)/usr/bin/

	# If you later add a config or procd service:
	# $(INSTALL_DIR) $(1)/etc/init.d
	# $(INSTALL_BIN) ./files/pttcomm.init $(1)/etc/init.d/pttcomm
endef

$(eval $(call BuildPackage,pttcomm))
