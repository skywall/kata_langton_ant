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
	var retColor Color
	if actualColor == White {
		retColor = Black
	} else {
		retColor = White
	}

	if actualAnt.direction == N && actualColor == White {
		return Ant{
			actualAnt.x + 1,
			actualAnt.y,
			E,
		}, retColor
	}
	if actualAnt.direction == N && actualColor == Black {
		return Ant{
			actualAnt.x - 1,
			actualAnt.y,
			W,
		}, retColor
	}
	if actualAnt.direction == S && actualColor == White {
		return Ant{
			actualAnt.x - 1,
			actualAnt.y,
			W,
		}, retColor
	}
	if actualAnt.direction == S && actualColor == Black {
		return Ant{
			actualAnt.x + 1,
			actualAnt.y,
			E,
		}, retColor
	}
	if actualAnt.direction == W && actualColor == White {
		return Ant{
			actualAnt.x,
			actualAnt.y - 1,
			N,
		}, retColor
	}

	if actualAnt.direction == W && actualColor == Black {
		return Ant{
			actualAnt.x,
			actualAnt.y + 1,
			S,
		}, retColor
	}
	if actualAnt.direction == E && actualColor == White {
		return Ant{
			actualAnt.x,
			actualAnt.y + 1,
			S,
		}, retColor
	}
	if actualAnt.direction == E && actualColor == Black {
		return Ant{
			actualAnt.x,
			actualAnt.y - 1,
			N,
		}, retColor
	}
	return Ant{3, 2, E}, Black
}
