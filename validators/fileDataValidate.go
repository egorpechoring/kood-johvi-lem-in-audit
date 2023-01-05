package validators

import (
	"main/errors"
	"strconv"
)

func ToInteger(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		errors.HandleError(errors.InvalidData)
	}
	return i
}
