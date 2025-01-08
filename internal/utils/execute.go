package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCmd(cmdName string, args ...string) string {
	cmd := exec.Command(cmdName, args...)

	fmt.Println(cmd)

	stdout, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return string(stdout)
}

func ExecuteNpmCmd(workingDir string, args ...string) string {
	arguments := []string{
		"--prefix",
		workingDir,
		"-y",
	}

	arguments = append(args, arguments...)

	return ExecuteCmd("npm", arguments...)
}

func ExecuteCmdInDocker(workingDir string, containerName string, args ...string) string {
	arguments := []string{
		"docker",
		"exec",
		"-w",
		workingDir,
		containerName,
	}

	arguments = append(arguments, args...)

	return ExecuteCmd("sudo", arguments...)
}
