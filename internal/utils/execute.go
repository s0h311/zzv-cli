package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCmd(cmdName string, args ...string) {
	cmd := exec.Command(cmdName, args...)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

func ExecuteNpmCmd(workingDir string, args ...string) {
	arguments := []string{
		"npm",
		"--prefix",
		workingDir,
	}

	arguments = append(arguments, args...)
}

func ExecuteCmdInDocker(workingDir string, args ...string) {
	arguments := []string{
		"exec",
		"-w",
		workingDir,
	}

	arguments = append(arguments, args...)

	ExecuteCmd("docker", arguments...)
}
