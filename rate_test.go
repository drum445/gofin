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

func TestRateWithLimit(t *testing.T) {
	_, err := RateWithLimit(14, 3.54017, 44.4057, 0, 0, 100)
	if err == nil {
		t.Errorf("limit was not triggered")
	}
}
