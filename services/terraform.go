package services

import (
	"os"
	"os/exec"
)

func TerraformInit(dir string, uuid string) error {

	cmd := exec.Command(
		"terraform",
		"init",
	)

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func TerraformApply(dir string) error {

	cmd := exec.Command(
		"terraform",
		"apply",
		"-auto-approve",
	)

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func TerraformDestroy(dir string) error {

	cmd := exec.Command(
		"terraform",
		"destroy",
		"-auto-approve",
	)

	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}