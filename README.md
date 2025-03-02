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

Run the `./scripts/morse_setup.sh` script to configure the build for your board of choice. Custom boards can be added to the `boards` folder and used as a target. See an existing board for the information which should be contained in the diffconfigs.

For example, to build for an EKH01
```
> ./scripts/morse_setup.sh -i -b ekh01
```

Target configuration files provided by this repository include

| Board       | Target                    |
|-------------|-------------------------- |
| EKH03v3     | `ekh03v3`                 |
| EKH03v4     | `ekh03v4`                 |
| EKH01v1     | `ekh01v1`                 |
| EKH01v2     | `ekh01v2`                 |
| EKH01-03    | `ekh01-03`                |
| EKH01       | `ekh01`                   |
| HL1         | `artini`                  |


After configuration is complete, run the build with
```
> make -j8
```

For verbose compilation, consider using
```
> make -j8 V=sc 2>&1 | tee log.txt
```

Once the build is complete a compiled image can be found in `bin/target/<platform>/<target>/`
