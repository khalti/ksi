package cmd

import(
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}