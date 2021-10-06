package gofin

import "math"

func FV(rate, nper, pmt, pv, loanType float64) float64 {
	if rate == 0.00 {
		return -1 * (pv + pmt*nper)
	}

	ratePerAnnum := rate
	pow := math.Pow((1 + ratePerAnnum), nper)
	fv := (pmt * (1 + ratePerAnnum*loanType) * (1 - pow) / ratePerAnnum) - pv*pow
	return fv
}
