package util

import (
	"os"
	"os/exec"
)

// Either returns the current user from global config
func GithubUser() string {
	gitUserCmd := exec.Command("git config --global user.name")
	output, _ := gitUserCmd.Output()

	return string(output)
}
