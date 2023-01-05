package readers

import (
	"bufio"
	"main/errors"
	"os"
)

func ReadFileLines(fileName string) []string {
	lines := []string{}

	file, err := os.Open("files/" + fileName)
	if err != nil {
		errors.HandleError(errors.FileOpenError)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		errors.HandleError(errors.ScannerError)
	}

	return lines
}
