# sbc-raspberrypi

This repo provides the overlay for building Raspberry Pi generic (models before Pi 5),
Raspberry Pi 5, and RevolutionPi Talos images.

For example usage instructions follow the [boot assets guide](https://www.talos.dev/latest/talos-guides/install/boot-assets/#example-raspberry-pi-overlay-with-imager).

## Overlay Options

| Option            | Description                                                         |
| ----------------- | ------------------------------------------------------------------- |
| `configTxt`       | Completely replace the `config.txt` file with the contents provided |
| `configTxtAppend` | Append the contents provided to the `config.txt` file               |
