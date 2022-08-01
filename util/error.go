package util

import(
	"os"
)

func HandleError(err error) {
	if NotNil(err) {
		Print(err.Error())
		os.Exit(1)
	}
}