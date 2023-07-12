package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func runFormatter(inputCode string, ch chan string) ([]byte, error) {
	cmd := exec.Command("rustfmt")
	cmd.Stdin = strings.NewReader(inputCode)
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut

	err := cmd.Run()
	if err != nil {
		if stdErr, ok := err.(*exec.ExitError); ok {
			ch <- ""
			return nil, fmt.Errorf("%s: %s", err.Error(), stdErr.String())
		}
		return nil, fmt.Errorf("Something went wrong: %s", err.Error())
	}
	ch <- string(stdOut.Bytes())
	return stdOut.Bytes(), nil
}
