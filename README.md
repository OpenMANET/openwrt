# MorseMicro OpenWrt
## Dependencies

To build the Morse Micro OpenWrt, you need a working Linux environment. This has been tested with Ubuntu 20.04 and higher.

Install build environment packages with
```
> sudo apt update
> sudo apt install build-essential clang flex g++ gawk gcc-multilib git gettext \
  libncurses5-dev libssl-dev python3-distutils rsync unzip zlib1g-dev swig
```

## Usage

Run the `./scripts/morse_setup.sh` script to configure the build for your board of choice.

For example, Using seeedstudio's WiFi Halow Modules on Raspberry Pi.
```
> ./scripts/morse_setup.sh -i -b ekh01
```

After configuration is complete, run the build with
```
> make -j8
```

For verbose compilation, consider using
```
> make -j8 V=sc 2>&1 | tee log.txt
```

Once the build is complete a compiled image can be found in `bin/target/<platform>/<target>/`

Of course, you can also download directly from the [release](https://github.com/Wvirfil123/openwrt/releases). The firmware in the release uses the [bcf_mf16858_fgh100mh_v6.3.0.bin](./bcf/bcf_mf16858_fgh100mh_v6.3.0.bin) BCF file.
  
If you compiled it yourself, after flashing it onto the Raspberry Pi, you need to SSH into the device and link `bcf_default.bin` to the `bcf_mf16858_fgh100mh_v6.3.0.bin` file.

```
cd /lib/firmware/morse
rm bcf_default.bin
ln -s bcf_mf16858_fgh100mh_v6.3.0.bin bcf_default.bin
```
