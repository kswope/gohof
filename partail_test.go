package gohof_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kswope/gohof"
)

func TestPartial(t *testing.T) {

	tests := []struct {
		x, y int
		want int
	}{
		{1, 1, 2},
		{2, 4, 6},
		{100, 200, 300},
	}

	for _, v := range tests {

		testName := fmt.Sprintf("%d+%d", v.x, v.y)

		t.Run(testName, func(t *testing.T) {

			adder := func(x, y int) int { return x + y }
			partialAdder := gohof.Partial1(adder, v.x)
			partialResult := partialAdder(v.y)

			if diff := cmp.Diff(partialResult, v.want); diff != "" {
				errorWithDiff(t, diff)
			}

		})

	}

}
