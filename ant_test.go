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
		testCase{},
	}

	for _, tc := range testCases {
		gotBoard, gotAnt := Step(tc.actualBoard, tc.actualAnt)
		if gotBoard != tc.nextBoard || gotAnt != tc.nextAnt {
			t.Errorf("Step(%v, %v) = %v, %v want %v, %v",
				tc.actualBoard, tc.actualAnt, gotBoard, gotAnt, tc.nextBoard, tc.nextAnt)
		}
	}
}
