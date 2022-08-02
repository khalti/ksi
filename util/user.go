package util

import (
	"fmt"
)

// Either returns the current user from global config
func GithubUser() string {
	output := Trim(Run("git", "config", "--global", "user.name"), " ", "\n")

	if output == ""{
		name := Prompt(`You don't have a user name configured yet.
Enter your Github username: `)
		Run("git","config", "--global", "user.name", fmt.Sprintf("%v", name))
		return name
	}

	return output
}
