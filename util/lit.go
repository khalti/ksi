package util

import(
	"fmt"
)

// Prints formatted strings
func Print(format string, a ...any)  {
	fmt.Println(fmt.Sprintf(format, a...))
}