package ant

func Step(actualBoard Board, actualAnt Ant) (Board, Ant) {
	actualColor := actualBoard[actualAnt.y][actualAnt.x]
	nextAnt, nextColor := AntStep(actualAnt, actualColor)

	nextBoard := actualBoard
	nextBoard[actualAnt.y][actualAnt.x] = nextColor

	return nextBoard, nextAnt
}

var nextDirectionTable = [][]Direction{
	/*   White||Black   */
	/* N */ {E, W},
	/* E */ {S, N},
	/* S */ {W, E},
	/* W */ {N, S},
}

var nextXoffset = [][]int{
	/*   White||Black   */
	/* N */ {1, -1},
	/* E */ {0, 0},
	/* S */ {-1, 1},
	/* W */ {0, 0},
}

var nextYoffset = [][]int{
	/*   White||Black   */
	/* N */ {0, 0},
	/* E */ {1, -1},
	/* S */ {0, 0},
	/* W */ {-1, 1},
}

var nextColorTable = [...]Color{Black, White}

func AntStep(actualAnt Ant, actualColor Color) (Ant, Color) {
	nextColor := nextColorTable[actualColor]
	nextDirection := nextDirectionTable[actualAnt.direction][actualColor]
	nextXoffset := nextXoffset[actualAnt.direction][actualColor]
	nextYoffset := nextYoffset[actualAnt.direction][actualColor]

	return Ant{
		actualAnt.x + nextXoffset,
		actualAnt.y + nextYoffset,
		nextDirection,
	}, nextColor
}
