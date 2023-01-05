package entities

type Room struct {
	Name      string
	X, Y      int
	IsVisited bool
	TunnelsTo []*Room
}
