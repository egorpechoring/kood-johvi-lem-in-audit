package builders

import (
	"main/entities"
	"main/errors"
	"main/validators"
	"strings"
)

func BuildAnthill(fileLines []string) (entities.Anthill, int) {
	anthill, l, isStart, isEnd := entities.Anthill{Rooms: map[string]*entities.Room{}}, len(fileLines), false, false

	n := validators.ToInteger(fileLines[0])
	if n == 0 {
		errors.HandleError(errors.InvalidData)
	}
	for i := 1; i < l; i++ {
		if fileLines[i] == "##start" {
			if i+1 < l {
				addMilestone(&anthill.StartName, fileLines[i+1])
				addRoom(&anthill, lineToRoom(fileLines[i+1]))
				i++
				isStart = true
			} else {
				errors.HandleError(errors.InvalidData)
			}
		} else if fileLines[i] == "##end" {
			if i+1 < l {
				addMilestone(&anthill.EndName, fileLines[i+1])
				addRoom(&anthill, lineToRoom(fileLines[i+1]))
				i++
				isEnd = true
			} else {
				errors.HandleError(errors.InvalidData)
			}
		} else if len(strings.Split(fileLines[i], " ")) == 3 {
			addRoom(&anthill, lineToRoom(fileLines[i]))
		} else if len(strings.Split(fileLines[i], "-")) == 2 {
			addTunnel(&anthill, fileLines[i])
		} else {
			if string(fileLines[i][0]) != "#" {
				errors.HandleError(errors.InvalidData)
			}
		}
	}
	if !(isStart && isEnd) {
		errors.HandleError(errors.InvalidData)
	}
	return anthill, n
}
