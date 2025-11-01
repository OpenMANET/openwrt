include $(TOPDIR)/rules.mk

PKG_NAME:=openmanet
PKG_VERSION:=0.0.1
PKG_RELEASE:=1

PKG_SOURCE:=$(PKG_NAME)-$(PKG_VERSION).tar.gz
PKG_SOURCE_URL:=https://codeload.github.com/openmanet/manager/tar.gz/$(PKG_VERSION)?
PKG_SOURCE_VERSION:=$(PKG_VERSION)
PKG_HASH:=skip

PKG_LICENSE:=MIT
PKG_MAINTAINER:=Corey Wagehoft

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_BUILD_FLAGS:=no-mips16

GO_PKG:=github.com/openmanet/manager
# GO_PKG_EXCLUDES:=test
GO_PKG_LDFLAGS_X:=cmd.Version=$(PKG_VERSION)

include $(INCLUDE_DIR)/package.mk
include ../../feeds/packages/lang/golang/golang-package.mk

define Package/openmanet
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=OpenMANET Manager
  DEPENDS:=$(GO_ARCH_DEPENDS) +kmod-batman-adv +alfred
endef

define Package/openmanet/description
A network management application for OpenMANET written in Go.
endef

define Package/openmanet/conffiles
/etc/openmanet/
endef

define Package/openmanet/install
	$(call GoPackage/Package/Install/Bin,$(1))

	$(INSTALL_DIR) $(1)/etc/openmanet/
	$(INSTALL_CONF) $(CURDIR)/files/sample_config.yml $(1)/etc/openmanet/config.yml
	$(INSTALL_DIR) $(1)/etc/config/
	$(INSTALL_CONF) $(CURDIR)/files/openmanet.config $(1)/etc/config/openmanet
	$(INSTALL_DIR) $(1)/etc/init.d/
	$(INSTALL_BIN) $(CURDIR)/files/openmanet.init $(1)/etc/init.d/openmanet
endef

$(eval $(call GoBinPackage,openmanet))
$(eval $(call BuildPackage,openmanet))
