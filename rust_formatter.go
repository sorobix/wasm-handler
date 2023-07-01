package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func runFormatter(inputCode string) ([]byte, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	cmd := exec.Command("rustfmt")
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
