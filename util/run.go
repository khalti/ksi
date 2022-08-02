package util

import "os/exec"

// Executes a shell command and returns the output as a string
func Run(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	stdOut, err := cmd.Output()

	HandleError(err)

	return string(stdOut)
}
