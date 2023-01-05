package entities

type Anthill struct {
	StartName, EndName string
	Rooms              map[string]*Room
}
