package ant

import "testing"

func TestStep(t *testing.T) {
	type testCase struct {
		actualBoard Board
		actualAnt   Ant
		nextBoard   Board
		nextAnt     Ant
	}

	testCases := []testCase{
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{2, 2, N},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{3, 2, E},
		},
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{3, 2, E},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{3, 3, S},
		},
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{3, 3, S},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, Black, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{2, 3, W},
		},
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, Black, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{2, 3, W},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{2, 2, N},
		},
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{2, 2, N},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, Black, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{1, 2, W},
		},
		testCase{
			actualBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, Black, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
			},
			actualAnt: Ant{1, 2, W},
			nextBoard: Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, Black, White, Black, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
			},
			nextAnt: Ant{1, 1, N},
		},
	}

	for _, tc := range testCases {
		gotBoard, gotAnt := Step(tc.actualBoard, tc.actualAnt)
		if gotBoard != tc.nextBoard || gotAnt != tc.nextAnt {
			t.Errorf("Step(\n%v%v\n) =\n%v%v \n\nwant\n\n%v%v",
				tc.actualBoard, tc.actualAnt, gotBoard, gotAnt, tc.nextBoard, tc.nextAnt)
		}
	}
}

func TestAntStep(t *testing.T) {
	type testCase struct {
		actualAnt   Ant
		actualColor Color
		nextAnt     Ant
		nextColor   Color
	}

	testCases := []testCase{
		testCase{Ant{2, 2, N}, White, Ant{3, 2, E}, Black},
		testCase{Ant{3, 2, E}, White, Ant{3, 3, S}, Black},
		testCase{Ant{3, 3, S}, White, Ant{2, 3, W}, Black},
		testCase{Ant{2, 3, W}, White, Ant{2, 2, N}, Black},
		testCase{Ant{1, 1, N}, White, Ant{2, 1, E}, Black},
		testCase{Ant{4, 4, S}, White, Ant{3, 4, W}, Black},
		testCase{Ant{4, 4, W}, White, Ant{4, 3, N}, Black},
		testCase{Ant{3, 1, E}, White, Ant{3, 2, S}, Black},
		testCase{Ant{2, 2, N}, Black, Ant{1, 2, W}, White},
		testCase{Ant{4, 4, S}, Black, Ant{5, 4, E}, White},
		testCase{Ant{4, 4, W}, Black, Ant{4, 5, S}, White},
		testCase{Ant{3, 1, E}, Black, Ant{3, 0, N}, White},
	}

	for _, tc := range testCases {
		gotAnt, gotColor := AntStep(tc.actualAnt, tc.actualColor)
		if gotAnt != tc.nextAnt || gotColor != tc.nextColor {
			t.Errorf("StepAnt(%v, %v) = %v, %v want %v, %v",
				tc.actualAnt, tc.actualColor, gotAnt, gotColor, tc.nextAnt, tc.nextColor)
		}
	}
}
