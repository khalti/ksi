package util

import (
	"fmt"
	"os/exec"
)

// Either returns the current user from global config
func GithubUser() string {
	gitUserCmd := exec.Command("git", "config", " --global", "user.name")
	output, err:= gitUserCmd.Output()

	if NotNil(err) || string(output) == ""{
		name := Prompt("You don't have a user name configured yet. What's your Github username?")
		setUserCmd := exec.Command("git", fmt.Sprintf("config --global user.name '%v'", name))

		// Ignore the output and/or the error
		std, err := setUserCmd.Output()
		if err != nil {
			Print(err.Error())
		}

		Print(string(std))

		return name
	}

	return string(output)
}
