package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
