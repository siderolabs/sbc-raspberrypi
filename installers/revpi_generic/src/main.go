// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

//go:embed config.txt
var configTxt []byte

func main() {
	adapter.Execute(&RpiInstaller{})
}

type RpiInstaller struct{}

type rpiOptions struct {
	ConfigTxt       string `yaml:"configTxt,omitempty"`
	ConfigTxtAppend string `yaml:"configTxtAppend,omitempty"`
}

func (i *RpiInstaller) GetOptions(extra rpiOptions) (overlay.Options, error) {
	return overlay.Options{
		Name: "revpi_generic",
		KernelArgs: []string{
			"console=tty0",
			"console=ttyAMA0,115200",
			"sysctl.kernel.kexec_load_disabled=1",
			"talos.dashboard.disabled=1",
		},
	}, nil
}

func (i *RpiInstaller) Install(options overlay.InstallOptions[rpiOptions]) error {
	err := copy.Dir(filepath.Join(options.ArtifactsPath, "arm64/firmware/boot"), filepath.Join(options.MountPrefix, "/boot/EFI"))
	if err != nil {
		return err
	}

	err = copy.Dir(filepath.Join(options.ArtifactsPath, "arm64/firmware/revpi_generic/boot"), filepath.Join(options.MountPrefix, "/boot/EFI"))
	if err != nil {
		return err
	}

	err = copy.File(filepath.Join(options.ArtifactsPath, "arm64/u-boot/rpi_generic/u-boot.bin"), filepath.Join(options.MountPrefix, "/boot/EFI/u-boot.bin"))
	if err != nil {
		return err
	}

	if options.ExtraOptions.ConfigTxt != "" {
		configTxt = []byte(options.ExtraOptions.ConfigTxt)
	}

	configTxt = append(configTxt, []byte("\n"+options.ExtraOptions.ConfigTxtAppend)...)

	return os.WriteFile(filepath.Join(options.MountPrefix, "/boot/EFI/config.txt"), configTxt, 0o644)
}
