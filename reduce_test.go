package gohof_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kswope/gohof"
)

func TestReduceStructs(t *testing.T) {

	type person struct {
		Name   string
		Weight int
	}

	people := []person{
		{"Kevin", 250},
		{"Bob", 200},
		{"Peg", 120},
		{"Chris", 280},
	}

	combine := func(accum int, p person) int {
		accum = accum + p.Weight
		return accum
	}

	totalWeight := gohof.Reduce(people, combine, 0)

	if diff := cmp.Diff(850, totalWeight); diff != "" {
		errorWithDiff(t, diff)
	}

}
