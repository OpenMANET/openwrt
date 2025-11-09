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

Run this to download all dependencies before starting a build.  It will make building more reliable.
```
> make download
```
Some targets provided by this repository include

| Board              | Target                    |
|--------------------|-------------------------- |
| MMx108-EKH01 (SDIO)| `mmx108-ekh01-sdio`       |
| MM6108-EKH01 (SPI) | `mm6108-ekh01-spi`        |
| MM8108-EKH01 (SPI) | `mm8108-ekh01-spi`        |
| HaLowLink1         | `halowlink1`              |
| MM8108-EKH19       | `ekh19`                   |

See the `boards/` folder for other targets and families.


After configuration is complete, run the build with
```
> make -j8
```

For verbose compilation, consider using
```
> make -j8 V=sc 2>&1 | tee log.txt
```

Once the build is complete a compiled image can be found in `bin/target/<platform>/<target>/`
