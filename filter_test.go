package gohof_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kswope/gohof"
)

func TestFilter(t *testing.T) {

	colors := []string{"blue", "red", "yellow", "green"}
	var expected []string
	var results []string

	results = gohof.Filter(colors, func(v string) bool { return len(v) > 3 })
	expected = []string{"blue", "yellow", "green"}
	if diff := cmp.Diff(expected, results); diff != "" {
		errorWithDiff(t, diff)
	}

	results = gohof.Filter(colors, func(v string) bool { return len(v) == 3 })
	expected = []string{"red"}
	if diff := cmp.Diff(expected, results); diff != "" {
		errorWithDiff(t, diff)
	}

}
