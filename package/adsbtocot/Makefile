include $(TOPDIR)/rules.mk

PKG_NAME:=adsbtocot
PKG_RELEASE:=1
PKG_MAINTAINER:=info@openmanet.net
PKG_LICENSE:=MIT
PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)
HOST_BUILD_DEPENDS:=python3/host

include $(INCLUDE_DIR)/package.mk

define Package/adsbtocot
  SECTION:=utils
  CATEGORY:=Utilities
  TITLE:=ADS-B to Cursor-on-Target (ADSBCOT) for OpenWrt (dump1090 local)
  DEPENDS:=+python3 +python3-pip +python3-setuptools +python3-wheel +dump1090
endef

define Package/adsbtocot/description
Installs ADSBCOT from PyPI and provides a procd service to forward ADS-B tracks
from dump1090 (Beast/SBS) to CoT/TAK (multicast or unicast).
endef

# Persist user edits
define Package/adsbtocot/conffiles
/etc/adsbcot/adsbcot.ini
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
endef

# Vendor ADSBCOT into the image root
define Build/Compile
	$(STAGING_DIR_HOSTPKG)/bin/pip3 install 'adsbcot[with_pymodes]' \
		--prefix=/usr --root='$(PKG_INSTALL_DIR)' --no-deps --no-cache-dir
endef

define Package/adsbtocot/install
	# Config + init
	$(INSTALL_DIR) $(1)/etc/adsbcot
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_CONF) ./files/adsbcot.ini  $(1)/etc/adsbcot/adsbcot.ini
	$(INSTALL_BIN)  ./files/adsbcot.init $(1)/etc/init.d/adsbcot

	# Python site-packages and console scripts
	$(INSTALL_DIR) $(1)/usr
	cp -a $(PKG_INSTALL_DIR)/usr/* $(1)/usr/

	# First-boot defaults to enable dump1090 network outputs and start ADSBCOT
	$(INSTALL_DIR) $(1)/etc/uci-defaults
	$(INSTALL_BIN) ./files/99-adsb-defaults $(1)/etc/uci-defaults/99-adsb-defaults
endef

$(eval $(call BuildPackage,adsbtocot))