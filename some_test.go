package gohof_test

import (
	"testing"

	"github.com/kswope/gohof"
)

func TestSome(t *testing.T) {

	ints := []int{1, 2, 3, 4, 5}

	pred := func(v int) bool {
		return v > 4
	}

	if !gohof.Some(ints, pred) {
		t.Error("should have found some")
	}

	pred = func(v int) bool {
		return v > 5
	}

	if gohof.Some(ints, pred) {
		t.Error("should not have found any")
	}

}
