package utils

import (
	"os/exec"
)

func ExecuteCmd(cmdName string, args ...string) string {
	cmd := exec.Command(cmdName, args...)

	stdout, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return string(stdout)
}

func ExecuteNpmCmd(workingDir string, args ...string) string {
	arguments := []string{
		"-c",
		"npm",
		"--prefix",
		workingDir,
		"-y",
	}

	arguments = append(args, arguments...)

	return ExecuteCmd("/bin/sh", arguments...)
}

func ExecuteCmdInDocker(workingDir string, containerName string, args ...string) string {
	arguments := []string{
		"-c",
		"docker",
		"exec",
		"-w",
		workingDir,
		containerName,
	}

	arguments = append(arguments, args...)

	return ExecuteCmd("/bin/sh", arguments...)
}
