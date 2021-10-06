package gofin

import (
	"testing"
)

func TestFV(t *testing.T) {
	fv1 := FV(5.8/100/12, 48, 2265.63, -133781.21, 0)
	fv1Expected := 46550.00026665449
	if fv1 != fv1Expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", fv1, fv1Expected)
	}

	fv2 := FV(5.8/100/12, 48, 2265.63, -133781.21, 1)
	fv2Expected := 45959.9959401056
	if fv2 != fv2Expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", fv2, fv2Expected)
	}
}
