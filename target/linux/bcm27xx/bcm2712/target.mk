#
# Copyright (C) 2017-2020 OpenWrt.org
#

include $(INCLUDE_DIR)/version.mk
include $(INCLUDE_DIR)/target.mk

SUBTARGET:=bcm2712
ARCH_PACKAGES:=

kmod-bcm27xx-rng_ARCH:=aarch64_cortex-a76
kmod-brcmfmac_ARCH:=aarch64_cortex-a76
kmod-brcmutil_ARCH:=aarch64_cortex-a76
kmod-mmc-bcm2835_ARCH:=aarch64_cortex-a76
kmod-rpi-eeprom_ARCH:=aarch64_cortex-a76

define Target/Devices
  define Device/raspberrypi
    DEVICE_VENDOR := Raspberry Pi
    DEVICE_TYPE := board
    DEVICE_CLASS := Single Board Computer
    DEVICE_SUBCLASS := BCM2712
    FEATURES += raspberrypi
    KERNEL_PATCHVER := 6.1
    KERNEL_INITRAMFS := Image
    KERNEL_NAME := $(KERNEL_INITRAMFS)
    KERNEL_LOADADDR := 0x80000
    DEVICE_DTS_DIR := $(DTS_DIR)/broadcom
    IMAGES += sysupgrade.img.gz
    IMAGE_KERNEL_OFFSET := 0x8000
    IMAGE_PART_BOOT_SIZE := 256
    IMAGE_PART_USER_SIZE := 2048
    IMAGE/sysupgrade.img.gz := sysupgrade-tar | gzip | pad-to 64k
  endef
endef
