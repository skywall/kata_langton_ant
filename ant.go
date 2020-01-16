package ant

func Step(actualBoard Board, actualAnt Ant) (Board, Ant) {
	if actualAnt.direction == E {
		return Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{3, 3, S}
	}
	return Board{
			{White, White, White, White, White},
			{White, White, White, White, White},
			{White, White, Black, White, White},
			{White, White, White, White, White},
			{White, White, White, White, White},
		},
		Ant{3, 2, E}
}

func AntStep(actualAnt Ant, actualColor Color) (Ant, Color) {
	retColor := [...]Color{Black, White}[actualColor]
	nextDirection := [][]Direction{
		/*   white||Black   */
		/* N */ {E, W},
		/* E */ {S, N},
		/* S */ {W, E},
		/* W */ {N, S},
	}[actualAnt.direction][actualColor]

	if actualAnt.direction == N && actualColor == White {
		return Ant{
			actualAnt.x + 1,
			actualAnt.y,
			nextDirection,
		}, retColor
	}
	if actualAnt.direction == N && actualColor == Black {
		return Ant{
			actualAnt.x - 1,
			actualAnt.y,
			nextDirection,
		}, retColor
	}
	if actualAnt.direction == S && actualColor == White {
		return Ant{
			actualAnt.x - 1,
			actualAnt.y,
			nextDirection,
		}, retColor
	}
	if actualAnt.direction == S && actualColor == Black {
		return Ant{
			actualAnt.x + 1,
			actualAnt.y,
			nextDirection,
		}, retColor
	}

	if actualAnt.direction == W && actualColor == White {
		return Ant{
			actualAnt.x,
			actualAnt.y - 1,
			nextDirection,
		}, retColor
	}

	if actualAnt.direction == W && actualColor == Black {
		return Ant{
			actualAnt.x,
			actualAnt.y + 1,
			nextDirection,
		}, retColor
	}
	if actualAnt.direction == E && actualColor == White {
		return Ant{
			actualAnt.x,
			actualAnt.y + 1,
			nextDirection,
		}, retColor
	}
	if actualAnt.direction == E && actualColor == Black {
		return Ant{
			actualAnt.x,
			actualAnt.y - 1,
			nextDirection,
		}, retColor
	}
	panic("Undefined step")
}
