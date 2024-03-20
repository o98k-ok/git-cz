package pkg

import (
	"bytes"
	"os/exec"
)

func ExecCommand(command string) (string, string, error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stderr = &errOut
	cmd.Stdout = &out

	err := cmd.Run()
	return out.String(), errOut.String(), err
}
