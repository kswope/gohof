package gohof_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kswope/gohof"
)

func TestFind(t *testing.T) {

	type fruit struct {
		Name  string
		Color string
	}

	fruits := []fruit{
		{"orange", "orange"},
		{"banana", "yellow"},
		{"pear", "yellow"},
		{"cherry", "red"},
	}

	f, found := gohof.Find(fruits, func(v fruit) bool { return v.Color == "yellow" })
	if !found {
		t.Fatal("Should have been found")
	}

	// found first in slice
	if diff := cmp.Diff(fruit{"banana", "yellow"}, f); diff != "" {
		errorWithDiff(t, diff)
	}

	_, found = gohof.Find(fruits, func(v fruit) bool { return v.Color == "purple" })
	if found {
		t.Fatal("Should not have been found")
	}

}
