package builders

import (
	"main/entities"
	"main/errors"
	"main/validators"
	"strings"
)

func lineToRoom(line string) *entities.Room {
	data := strings.Split(line, " ")
	if len(data) != 3 {
		errors.HandleError(errors.InvalidData)
	}
	return &entities.Room{
		Name:      data[0],
		X:         validators.ToInteger(data[1]),
		Y:         validators.ToInteger(data[2]),
		TunnelsTo: [](*entities.Room){},
		IsVisited: false,
	}
}

func addMilestone(rn *string, str string) {
	room := lineToRoom(str)
	*rn = room.Name
}

func addRoom(anthill *entities.Anthill, r *entities.Room) {
	_, found := (*anthill).Rooms[(*r).Name]
	if found {
		errors.HandleError(errors.SameNameError)
	}
	(*anthill).Rooms[r.Name] = r
}

func addTunnel(anthill *entities.Anthill, line string) {
	data := strings.Split(line, "-")
	if len(data) != 2 {
		errors.HandleError(errors.InvalidData)
	}
	r1, err1 := (*&anthill.Rooms)[data[0]]
	r2, err2 := (*&anthill.Rooms)[data[1]]
	if !err1 || !err2 {
		errors.HandleError(errors.InvalidData)
	}
	if r1.Name == r2.Name {
		return
	}
	r1.TunnelsTo = append(r1.TunnelsTo, r2)
	r2.TunnelsTo = append(r2.TunnelsTo, r1)
}
