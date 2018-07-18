package finance

import "math"

func NPER(rate, pmt, pv, fv, loanType float64) float64 {
	if rate == 0.00 {
		nper := (-(pv + fv) / pmt)
		return math.Round(nper)
	}

	ratePerAnnum := (rate / 100) / 12
	num := pmt*(1+ratePerAnnum*loanType) - fv*ratePerAnnum
	den := (pv*ratePerAnnum + pmt*(1+ratePerAnnum*loanType))
	nper := math.Log10(num/den) / math.Log10(1+ratePerAnnum)
	return math.Round(nper)
}
