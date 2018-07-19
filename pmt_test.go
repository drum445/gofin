package gofin

import (
	"testing"
)

func TestPMT(t *testing.T) {
	pmt := PMT(9.3, 36, -12995, 6000, 0)
	pmtExpected := 269.9170849608034
	if pmt != pmtExpected {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", pmt, pmtExpected)
	}
}
