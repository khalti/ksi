package util

import (
	"os/exec"
)

// Either returns the current user from global config
func GithubUser() string {
	gitUserCmd := exec.Command("git config --global user.name")
	output, err:= gitUserCmd.Output()

	if NotNil(err) {
		
	}

	return string(output)
}
