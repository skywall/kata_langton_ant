package ant

type Color int

const (
	White = iota
	Black
)

type Direction int

const (
	N = iota
	E
	S
	W
)

type Board [5][5]Color

type Ant struct {
	x         int
	y         int
	direction Direction
}

func Step(actualBoard Board, actualAnt Ant) (Board, Ant) {
	return Board{}, Ant{}
}
