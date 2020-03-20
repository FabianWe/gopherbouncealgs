package tests

import (
	"github.com/FabianWe/gopherbouncealgs"
	"testing"
)

// TestPow tests utils.Pow.
func TestPow(t *testing.T) {
	tests := []struct {
		base, exp, out int64
	}{
		{42, 0, 1},
		{0, 42, 0},
		{2, 8, 256},
		{10, 10, 10000000000},
	}
	for _, tc := range tests {
		if got := gopherbouncealgs.Pow(tc.base, tc.exp); got != tc.out {
			t.Errorf("Expected that %d ^ %d = %d, got %d", tc.base, tc.exp, tc.out, got)
		}
	}
}
