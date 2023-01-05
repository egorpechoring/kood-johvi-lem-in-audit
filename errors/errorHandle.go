package errors

import (
	"fmt"
	"os"
)

func HandleError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
