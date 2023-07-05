package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func runFormatter(inputCode string) ([]byte, error) {
	cmd := exec.CommandContext(context.Background(), "rustfmt")
	cmd.Stdin = strings.NewReader(inputCode)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut

	err := cmd.Run()
	if err != nil {
		if stdErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("%s: %s", err.Error(), stdErr.String())
		}
		return nil, fmt.Errorf("Something went wrong: %s", err.Error())
	}

	return stdOut.Bytes(), nil
}
