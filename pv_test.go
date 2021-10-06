package gofin

import (
	"testing"
)

func TestPV(t *testing.T) {
	pv := PV(9.3/100/12, 36, 350.50, 1000, 0)
	pvExpected := -11731.212657351734
	if pv != pvExpected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", pv, pvExpected)
	}
}
