package builders

import (
	"main/entities"
	"main/errors"
)

func BuildRoutes(anthill entities.Anthill) [][]string {
	allRoutes := [][]string{}
	recursiveBuildRoutes(&allRoutes, anthill, []string{}, anthill.StartName)
	if len(allRoutes) == 0 {
		errors.HandleError(errors.InvalidData)
	}
	return allRoutes
}

func recursiveBuildRoutes(allRoutes *[][]string, anthill entities.Anthill, path []string, targetRoomName string) {
	pathCopy := make([]string, len(path))
	copy(pathCopy, path)

	currentRoom, _ := anthill.Rooms[targetRoomName]
	contained := false
	for _, s := range pathCopy {
		if s == targetRoomName {
			contained = true
			break
		}
	}

	if contained {
		return
	}

	pathCopy = append(pathCopy, currentRoom.Name)

	if currentRoom.Name == anthill.EndName {
		*allRoutes = append(*allRoutes, pathCopy)
		return
	}

	for _, targetRoom := range currentRoom.TunnelsTo {
		if !targetRoom.IsVisited {
			recursiveBuildRoutes(allRoutes, anthill, pathCopy, targetRoom.Name)
		}
	}
}
