# sbc-raspberrypi

This repo provides the overlay for Raspberry Pi generic Talos image.

For example usage instructions follow the [boot assets guide](https://www.talos.dev/latest/talos-guides/install/boot-assets/#example-raspberry-pi-overlay-with-imager).

## Overlay Options

| Option            | Description                                                         |
| ----------------- | ------------------------------------------------------------------- |
| `configTxt`       | Completely replace the `config.txt` file with the contents provided |
| `configTxtAppend` | Append the contents provided to the `config.txt` file               |
