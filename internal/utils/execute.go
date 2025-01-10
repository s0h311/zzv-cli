package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCmd(cmdName string, args ...string) string {
	normalizedArgs := []string{}

	for _, arg := range args {
		normalizedArgs = append(normalizedArgs, strings.Split(arg, " ")...)
	}

	cmd := exec.Command(cmdName, normalizedArgs...)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(string(stdout))
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
