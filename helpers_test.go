package gohof_test

import "testing"

func errorWithDiff(t *testing.T, diff any) {
	t.Errorf("mismatch (-want +got):\n%s", diff)
}
