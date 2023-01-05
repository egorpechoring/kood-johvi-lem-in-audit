package printers

import "fmt"

func PrintFileLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
