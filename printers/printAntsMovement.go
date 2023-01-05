package printers

import (
	"fmt"
	"strconv"
)

func PrintMovement(strategy [][]string, antsN, steps int) {
	data, a, t := []string{}, 1, 1

	for s := 0; s < steps; s++ {
		data = append(data, "")
	}

	for _, path := range strategy {
		n := steps - len(path) + 1
		for i := 0; i < n+1; i++ {
			if a > antsN {
				break
			}
			for r, room := range path {
				if room != strategy[0][0] {
					data[t-2+r] += "L" + strconv.Itoa(a) + "-" + room + " "
				}
			}
			a++
			t++
		}
		t = 1
	}

	for _, d := range data {
		fmt.Println(d)
	}
}
