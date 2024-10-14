# SPDX-License-Identifier: GPL-2.0-only
#
# Copyright (C) 2012-2019 OpenWrt.org
# Copyright (C) 2016-2017 LEDE project

include $(TOPDIR)/rules.mk
include $(INCLUDE_DIR)/image.mk

DEVICE_VARS += SYSINFO_MODEL SYSINFO_BOARD_NAME SPI_OVERLAY

FAT32_BLOCK_SIZE=1024
FAT32_BLOCKS=$(shell echo $$(($(CONFIG_TARGET_KERNEL_PARTSIZE)*1024*1024/$(FAT32_BLOCK_SIZE))))

define Build/Compile
	$(CP) $(LINUX_DIR)/COPYING $(KDIR)/COPYING.linux
endef

### Image scripts ###
define Build/boot-common
	rm -f $@.boot
	mkfs.fat -n boot -C $@.boot $(FAT32_BLOCKS)
	mcopy -i $@.boot $(KDIR)/COPYING.linux ::
	mcopy -i $@.boot $(KDIR)/LICENCE.broadcom ::
	mcopy -i $@.boot cmdline.txt ::
	mcopy -i $@.boot config.txt ::
	mcopy -i $@.boot distroconfig.txt ::
	mcopy -i $@.boot $(IMAGE_KERNEL) ::$(KERNEL_IMG)
	$(foreach dts,$(shell echo $(DEVICE_DTS)),mcopy -i $@.boot $(DTS_DIR)/$(dts).dtb ::;)
	mmd -i $@.boot ::/overlays
	mcopy -i $@.boot $(DTS_DIR)/overlays/*.dtbo ::/overlays/
	mcopy -i $@.boot $(DTS_DIR)/overlays/README ::/overlays/
endef

define Build/boot-2708
	mcopy -i $@.boot $(KDIR)/bootcode.bin ::
	mcopy -i $@.boot $(KDIR)/start.elf ::
	mcopy -i $@.boot $(KDIR)/start_cd.elf ::
	mcopy -i $@.boot $(KDIR)/start_x.elf ::
	mcopy -i $@.boot $(KDIR)/fixup.dat ::
	mcopy -i $@.boot $(KDIR)/fixup_cd.dat ::
	mcopy -i $@.boot $(KDIR)/fixup_x.dat ::
endef

define Build/boot-2711
	mcopy -i $@.boot $(KDIR)/start4.elf ::
	mcopy -i $@.boot $(KDIR)/start4cd.elf ::
	mcopy -i $@.boot $(KDIR)/start4x.elf ::
	mcopy -i $@.boot $(KDIR)/fixup4.dat ::
	mcopy -i $@.boot $(KDIR)/fixup4cd.dat ::
	mcopy -i $@.boot $(KDIR)/fixup4x.dat ::
endef

define Build/boot-ekh01
	# Put ekh01 model info into device tree (picked up by 01_sysinfo).
	# This uses a custom overlay which accepts parameters (via RPI's __overrides__ mechanism),
	# and writes those parameters into the root of the device tree (as sysinfo-model
	# and sysinfo-board-name).
	# Note that we cannot easily put it on the rootfs, because our multi-profile builds share this
	# (i.e. they only differ in /boot).
	# We can't easily use /boot directly because that's not mounted early in the preinit process.
	# To support both MM610x and MM810x on SPI, each requiring a different
	# device tree overlay, we have updated the SPI overlay configuration in
	# `distroconfig` to apply the appropriate overlay based on the device variant.

	echo -e '\ndtoverlay=sysinfo,board-name="$(SYSINFO_BOARD_NAME)",model="$(SYSINFO_MODEL)"\ndtoverlay=$(SPI_OVERLAY)' | cat boards/ekh01/distroconfig.txt - | mcopy -o -i $@.boot - ::/distroconfig.txt

endef

define Build/sdcard-img
	./gen_rpi_sdcard_img.sh $@ $@.boot $(IMAGE_ROOTFS) \
		$(CONFIG_TARGET_KERNEL_PARTSIZE) $(CONFIG_TARGET_ROOTFS_PARTSIZE)
endef

### Devices ###
define Device/Default
  PROFILES := Default
  DEVICE_VENDOR := Raspberry Pi
  KERNEL := kernel-bin
  KERNEL_IMG := kernel.img
  IMAGES := factory.img.gz sysupgrade.img.gz
  IMAGE/sysupgrade.img.gz := boot-common | boot-2708 | sdcard-img | gzip | append-metadata
  IMAGE/factory.img.gz := boot-common | boot-2708 | sdcard-img | gzip
endef

define Device/rpi
  DEVICE_MODEL := B/B+/CM/Zero/ZeroW
  DEVICE_DTS := \
	bcm2708-rpi-b bcm2708-rpi-b-rev1 bcm2708-rpi-b-plus \
	bcm2708-rpi-cm \
	bcm2708-rpi-zero bcm2708-rpi-zero-w
  SUPPORTED_DEVICES := \
	rpi-b rpi-b-plus rpi-cm rpi-zero rpi-zero-w \
	raspberrypi,model-b raspberrypi,model-b-plus raspberrypi,model-b-rev2 \
	raspberrypi,compute-module raspberrypi,compute-module-1 \
	raspberrypi,model-zero raspberrypi,model-zero-w
  DEVICE_PACKAGES := \
	cypress-firmware-43430-sdio \
	brcmfmac-nvram-43430-sdio \
	kmod-brcmfmac wpad-basic-mbedtls
endef
ifeq ($(SUBTARGET),bcm2708)
  TARGET_DEVICES += rpi
endif

define Device/rpi-2
  DEVICE_MODEL := 2B/2B 1.2
  DEVICE_VARIANT := (32bit)
  DEVICE_ALT0_VENDOR := Raspberry Pi
  DEVICE_ALT0_MODEL := 3B/3B+/CM3
  DEVICE_ALT0_VARIANT := (32bit)
  DEVICE_ALT1_VENDOR := Raspberry Pi
  DEVICE_ALT1_MODEL := 4B/400/CM4
  DEVICE_ALT1_VARIANT := (32bit)
  DEVICE_DTS := \
	bcm2709-rpi-2-b bcm2710-rpi-2-b \
	bcm2710-rpi-3-b bcm2710-rpi-3-b-plus \
	bcm2711-rpi-4-b bcm2711-rpi-400 \
	bcm2710-rpi-cm3 bcm2711-rpi-cm4 \
	bcm2710-rpi-zero-2
  SUPPORTED_DEVICES := \
	rpi-2-b rpi-3-b rpi-3-b-plus rpi-cm rpi-zero-2 \
	raspberrypi,2-model-b raspberrypi,2-model-b-rev2 \
	raspberrypi,3-model-b raspberrypi,3-model-b-plus \
	raspberrypi,3-compute-module raspberrypi,compute-module-3 \
	raspberrypi,400 raspberrypi,4-compute-module raspberrypi,4-model-b \
	raspberrypi,model-zero-2
  DEVICE_PACKAGES := \
	cypress-firmware-43430-sdio \
	brcmfmac-nvram-43430-sdio \
	cypress-firmware-43455-sdio \
	brcmfmac-nvram-43455-sdio \
	kmod-brcmfmac wpad-basic-mbedtls
  IMAGE/sysupgrade.img.gz := boot-common | boot-2708 | boot-2711 | sdcard-img | gzip | append-metadata
  IMAGE/factory.img.gz := boot-common | boot-2708 | boot-2711 | sdcard-img | gzip
endef
ifeq ($(SUBTARGET),bcm2709)
  TARGET_DEVICES += rpi-2
endif

define Device/rpi-3
  DEVICE_MODEL := 3B/3B+/CM3
  DEVICE_VARIANT := (64bit)
  DEVICE_ALT0_VENDOR := Raspberry Pi
  DEVICE_ALT0_MODEL := 2B-1.2
  DEVICE_ALT0_VARIANT := (64bit)
  KERNEL_IMG := kernel8.img
  DEVICE_DTS := \
	broadcom/bcm2710-rpi-2-b \
	broadcom/bcm2710-rpi-3-b broadcom/bcm2710-rpi-3-b-plus \
	broadcom/bcm2710-rpi-cm3 \
	broadcom/bcm2710-rpi-zero-2
  SUPPORTED_DEVICES := \
	rpi-3-b rpi-3-b-plus rpi-zero-2 \
	raspberrypi,2-model-b-rev2 \
	raspberrypi,3-model-b raspberrypi,3-model-b-plus \
	raspberrypi,3-compute-module raspberrypi,compute-module-3 \
	raspberrypi,model-zero-2
  DEVICE_PACKAGES := \
	cypress-firmware-43430-sdio \
	brcmfmac-nvram-43430-sdio \
	cypress-firmware-43455-sdio \
	brcmfmac-nvram-43455-sdio \
	kmod-brcmfmac wpad-basic-mbedtls
endef
ifeq ($(SUBTARGET),bcm2710)
  TARGET_DEVICES += rpi-3
endif

define Device/rpi-4
  DEVICE_MODEL := 4B/400/CM4
  DEVICE_VARIANT := (64bit)
  KERNEL_IMG := kernel8.img
  DEVICE_DTS := \
	broadcom/bcm2711-rpi-400 \
	broadcom/bcm2711-rpi-4-b \
	broadcom/bcm2711-rpi-cm4
  SUPPORTED_DEVICES := \
	raspberrypi,400 \
	raspberrypi,4-compute-module \
	raspberrypi,4-model-b
  DEVICE_PACKAGES := \
	cypress-firmware-43455-sdio \
	brcmfmac-nvram-43455-sdio \
	kmod-brcmfmac wpad-basic-mbedtls \
	kmod-usb-net-lan78xx \
	kmod-r8169
  IMAGE/sysupgrade.img.gz := boot-common | boot-2711 | sdcard-img | gzip | append-metadata
  IMAGE/factory.img.gz := boot-common | boot-2711 | sdcard-img | gzip
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += rpi-4
endif

define Device/morse_ekh01
  $(Device/rpi-4)
  DEVICE_VENDOR := MorseMicro
  DEVICE_MODEL := MM6108-EKH01
  DEVICE_VARIANT :=
  # remove the factory image
  IMAGES := sysupgrade.img.gz
  # Simplify by removing board and OpenWRT version.
  IMAGE_PREFIX = $$(VERSION_DIST_SANITIZED)-$$(IMG_PREFIX_VERCODE)$$(IMG_PREFIX_EXTRA)$$(call sanitize,$$(DEVICE_TITLE))
  DEVICE_IMG_NAME = $$(IMAGE_PREFIX)-$$(1)-$$(2)
  # This is very similar to DEVICE_TITLE, but we want to handle the 'variants' differently
  # (i.e. EKH01v1 vs EKH01 v1, and EKH01-03 vs EKH01 03).
  SYSINFO_MODEL := MorseMicro EKH01
  # SYSINFO_BOARD_NAME will usually be the first entry in SUPPORTED_DEVICES,
  # This is the name on device; SUPPORTED_DEVICES are the names in the metadata.
  # SYSINFO_BOARD_NAME must be one of the names in SUPPORTED_DEVICES for upgrades
  # to work.
  SYSINFO_BOARD_NAME := morse,ekh01
  SUPPORTED_DEVICES := morse,ekh01 morse,ekh01-05us
  SPI_OVERLAY := mm610x-spi
  DEVICE_PACKAGES += kmod-morse netifd-morse
  IMAGE/sysupgrade.img.gz := boot-common | boot-2711 | boot-ekh01 | sdcard-img | gzip | append-metadata
  IMAGE/factory.img.gz := boot-common | boot-2711 | boot-ekh01 | sdcard-img | gzip
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01
endif

define Device/morse_ekh01v1
  $(Device/morse_ekh01)
  DEVICE_VARIANT := v1
  SYSINFO_MODEL := MorseMicro EKH01v1
  SYSINFO_BOARD_NAME := morse,ekh01v1
  SUPPORTED_DEVICES := morse,ekh01v1 morse,ekh01
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01v1
endif

define Device/morse_ekh01v2
  $(Device/morse_ekh01)
  DEVICE_VARIANT := v2
  SYSINFO_MODEL := MorseMicro EKH01v2
  SYSINFO_BOARD_NAME := morse,ekh01v2
  SUPPORTED_DEVICES := morse,ekh01v2
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01v2
endif

define Device/morse_ekh01-03
  $(Device/morse_ekh01)
  DEVICE_VARIANT := 03
  SYSINFO_MODEL := MorseMicro EKH01-03
  SYSINFO_BOARD_NAME := morse,ekh01-03
  SUPPORTED_DEVICES := morse,ekh01-03
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01-03
endif

define Device/morse_ekh01-mf13455
  $(Device/morse_ekh01)
  DEVICE_MODEL := MM8108-EKH01
  DEVICE_VARIANT := mf13455
  SYSINFO_MODEL := MorseMicro EKH01-mf13455
  SYSINFO_BOARD_NAME := morse,ekh01-mf13455
  SUPPORTED_DEVICES := morse,ekh01-mf13455
  SPI_OVERLAY := mm810x-spi
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01-mf13455
endif

define Device/morse_ekh01-mf15457
  $(Device/morse_ekh01)
  DEVICE_MODEL := MM8108-EKH01
  DEVICE_VARIANT := mf15457
  SYSINFO_MODEL := MorseMicro EKH01-mf15457
  SYSINFO_BOARD_NAME := morse,ekh01-mf15457
  SUPPORTED_DEVICES := morse,ekh01-mf15457
endef
ifeq ($(SUBTARGET),bcm2711)
  TARGET_DEVICES += morse_ekh01-mf15457
endif

$(eval $(call BuildImage))
