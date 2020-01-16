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
			nextAnt: Ant{3, 2, N},
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
