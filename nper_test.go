package gofin

import (
	"testing"
)

func TestNPER(t *testing.T) {
	nper := NPER(10, 2625.73, -81374.62, 0, 0)
	nperExpected := 36.0000015398228
	if nper != nperExpected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", nper, nperExpected)
	}
}
