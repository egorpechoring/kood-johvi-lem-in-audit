package main

import (
	"fmt"
	"main/builders"
	"main/errors"
	"main/printers"
	"main/readers"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		errors.HandleError(errors.InputError)
	}
	fileLines := readers.ReadFileLines(os.Args[1])

	anthill, antsN := builders.BuildAnthill(fileLines)
	allRoutes := builders.BuildRoutes(anthill)
	bestStrategy, bestSteps := builders.BuildStrategy(anthill.StartName, anthill.EndName, allRoutes, antsN)

	printers.PrintFileLines(fileLines)
	fmt.Println()
	printers.PrintMovement(bestStrategy, antsN, bestSteps)

	// Uncomment to see selected strategy and steps amount
	// fmt.Println(bestStrategy)
	// fmt.Println(bestSteps)
}
