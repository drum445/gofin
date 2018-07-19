package gofin

import (
	"testing"
)

func TestRate(t *testing.T) {
	rate := Rate(60, 1600, -75304.59, 0, 0)
	rateExpected := 10.000000251301236
	if rate != rateExpected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", rate, rateExpected)
	}
}
