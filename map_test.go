package gohof_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kswope/gohof"
)

func TestMap(t *testing.T) {

	type account struct {
		Name    string
		Balance int
	}

	bank := []account{
		{"Kevin", 100},
		{"Jeff", 1_000_000_000},
	}

	// add a buck
	deposit := func(a account) account {
		a.Balance += 1
		return a
	}

	bank = gohof.Map(bank, deposit)

	expected := []account{{"Kevin", 101}, {"Jeff", 1_000_000_001}}
	if diff := cmp.Diff(expected, bank); diff != "" {
		errorWithDiff(t, diff)
	}

}
