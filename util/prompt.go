package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Asks a user for prompt
func Prompt(label string) string {
	fmt.Print(label)
	reader:= bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	HandleError(err)

	return strings.TrimSpace(line)
} 