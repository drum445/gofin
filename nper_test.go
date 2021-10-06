package gofin

import (
	"testing"
)

func TestNPER(t *testing.T) {
	nper := NPER(10/100/12, 2625.73, -81374.62, 0, 0)
	nperExpected := 30.99123672273996
	if nper != nperExpected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", nper, nperExpected)
	}
}
