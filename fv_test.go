package gofin

import (
	"testing"
)

func TestFV(t *testing.T) {
	fv1 := FV(5.8, 48, 2265.63, -133781.21, 0)
	fv1Expected := 46550.000266654475
	if fv1 != fv1Expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", fv1, fv1Expected)
	}
}
