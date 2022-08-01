package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Asks a user for prompt
func Prompt(label string) string {
	var s string
	r:= bufio.NewReader(os.Stdin)
	for{
		fmt.Fprint(os.Stderr, label+" ")
		s, err := r.ReadString('\n')

		HandleError(err)

		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
} 