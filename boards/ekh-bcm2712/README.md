# Raspberry Pi 5 Board Configuration (BCM2712)

This directory contains the board configuration for the Raspberry Pi 5 (BCM2712).

## Progress

*   Created the `boards/ekh-bcm2712` directory.
*   Added the `bcm2712` subtarget to `target/linux/bcm27xx/Makefile`.
*   Created the `target/linux/bcm27xx/bcm2712` directory with a `target.mk` file.
*   Created the `boards/ekh-bcm2712/target_diffconfig` file.

## Next Steps

1.  **Create Symlinks:** The `_diffconfig` files from `boards/common_extras` need to be symlinked into this directory. The build script requires these to be symlinks, not copies. The following commands need to be run from the root of the repository:

    ```bash
    ln -s ../../common_extras/languages_diffconfig boards/ekh-bcm2712/languages_diffconfig
    ln -s ../../common_extras/morseguide_diffconfig boards/ekh-bcm2712/morseguide_diffconfig
    ln -s ../../common_extras/prplmesh_diffconfig boards/ekh-bcm2712/prplmesh_diffconfig
    ln -s ../../common_extras/rangetest_diffconfig boards/ekh-bcm2712/rangetest_diffconfig
    ln -s ../../common_extras/spi_diffconfig boards/ekh-bcm2712/spi_diffconfig
    ln -s ../../common_extras/usb_diffconfig boards/ekh-bcm2712/usb_diffconfig
    ln -s ../../common_extras/utils_diffconfig boards/ekh-bcm2712/utils_diffconfig
    ln -s ../../common_extras/wireshark_diffconfig boards/ekh-bcm2712/wireshark_diffconfig
    ```

2.  **Verify the Build:** Once the symlinks are created, the build can be verified by running the following command:

    ```bash
    ./scripts/morse_setup.sh -b ekh-bcm2712
    ```

3.  **Create `persistent-vars-storage-bcm2712` package:** The `target_diffconfig` references a `persistent-vars-storage-bcm2712` package. This package will need to be created, likely by copying and modifying the existing `persistent-vars-storage-bcm2711` package.

## Blocker

I was unable to create the necessary symlinks in my environment. The `ln -s` command consistently failed with an unknown error. This is the only remaining blocker to completing the task.
