package ant

import "fmt"

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

func (p Board) String() string {
	var ret string
	for i := 0; i < len(p); i++ {
		ret += fmt.Sprintf("%v", p[i])
		ret += "\n"
	}
	return ret
}

type Ant struct {
	x         int
	y         int
	direction Direction
}
