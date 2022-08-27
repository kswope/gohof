package gohof_test

import (
	"strings"
	"testing"

	"github.com/kswope/gohof"
)

func TestEvery(t *testing.T) {

	var testStrings []string

	predicate := func(v string) bool {
		return strings.Contains(v, "test")
	}

	testStrings = []string{
		"test",
		"detest",
		"latest",
		"lutestring",
		"testosterone",
	}

	if gohof.Every(testStrings, predicate) == false {
		t.Error("failed")
	}

	testStrings = []string{
		"test",
		"quiz",
		"detest",
		"lutestring",
		"testosterone",
	}

	if gohof.Every(testStrings, predicate) == true {
		t.Error("failed")
	}

}
