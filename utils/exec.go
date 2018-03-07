package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func Popen(command ...string) (string, error) {
	var output bytes.Buffer
	var err error
	var args []string
	var bin string

	if len(command) == 0 {
		return "", errors.New("Must provide Exec with at least one argument")
	}

	if len(command) == 1 {
		bin = command[0]
	} else if len(command) >= 1 {
		bin = command[0]
		args = command[1:]
	}

	cmd := exec.Command(bin, args...)
	cmd.Stdout = &output
	err = cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprintf("%s happened...", err))
	}
	return output.String(), nil
}
