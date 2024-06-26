package core

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2024
	Description: The Vanilla System Operator is a package manager,
	a system updater and a task automator.
*/

import (
	"github.com/vanilla-os/apx/v2/core"
)

func GetPico() (*core.SubSystem, error) {
	subsystem, err := core.LoadSubSystem("vso-pico", false)
	if err != nil {
		return nil, err
	}

	return subsystem, nil
}

func PicoExists() bool {
	_, err := GetPico()
	return err == nil
}

func PicoInit() error {
	stack, _ := core.LoadStack("vso-pico")
	subsystem, err := core.NewSubSystem(
		"vso-pico",
		stack,
		"",
		true,
		true,
		false,
		false,
		true,
		"",
	)

	if err != nil {
		return err
	}

	err = subsystem.Create()
	return err
}

func PicoDelete() error {
	subsystem, err := GetPico()
	if err != nil {
		return err
	}

	err = subsystem.Remove()
	if err != nil {
		return err
	}

	return nil
}

func PicoExport(app string, binary string) error {
	subsystem, err := GetPico()
	if err != nil {
		return err
	}

	if app != "" {
		err = subsystem.ExportDesktopEntry(app)
	} else {
		err = subsystem.ExportBin(binary, "")
	}

	return err
}

func PicoUnexport(app string, binary string) error {
	subsystem, err := GetPico()
	if err != nil {
		return err
	}

	if app != "" {
		err = subsystem.UnexportDesktopEntry(app)
	} else {
		err = subsystem.UnexportBin(binary, "")
	}

	return err
}

func PicoUpgrade() error {
	pico, err := GetPico()
	if err != nil {
		return err
	}

	pkgManager, err := pico.Stack.GetPkgManager()
	if err != nil {
		return err
	}

	finalArgs := pkgManager.GenCmd(pkgManager.CmdUpdate, []string{}...)
	_, err = pico.Exec(false, false, finalArgs...)
	if err != nil {
		return err
	}

	finalArgs = pkgManager.GenCmd(pkgManager.CmdUpgrade, []string{}...)
	_, err = pico.Exec(false, false, finalArgs...)
	if err != nil {
		return err
	}

	return nil
}
