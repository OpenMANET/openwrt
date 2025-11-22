# MorseMicro OpenWrt
## Dependencies

To build the Morse Micro OpenWrt, you need a working Linux environment. This has been tested with Ubuntu 24.04.

Install build environment packages with
```
> sudo apt update
> sudo apt install build-essential clang flex bison g++ gawk \
gcc-multilib g++-multilib gettext git libncurses5-dev libssl-dev \
python3-setuptools rsync swig unzip zlib1g-dev file wget libnl-3-dev \
libnl-genl-3-dev libgps-dev libcap-dev pkg-config libopus-dev \
libopusfile-dev portaudio19-dev net-tools
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

After configuration is complete, run the build with
```
> make -j8
```

For verbose compilation, consider using
```
> make -j8 V=sc 2>&1 | tee log.txt
```

Once the build is complete a compiled image can be found in `bin/target/<platform>/<target>/`
