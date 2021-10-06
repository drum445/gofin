package gofin

import "math"

func PMT(rate, nper, pv, fv, loanType float64) float64 {
	pmt := 0.00
	pv = pv * -1
	pvMinusFV := pv - fv

	if rate == 0.00 {
		pmt = pvMinusFV / nper
	} else {
		raterPerAnnum := rate
		rateToNPER := math.Pow((raterPerAnnum + 1), nper)
		pmt = (pvMinusFV * (raterPerAnnum * rateToNPER)) / (rateToNPER - 1)

		fvRate := fv * raterPerAnnum
		pmt = (pmt + fvRate)

		// loan type is when the payment is being paid
		// 0 == end of month (default), 1 == start of month
		// so one less months worth of rate is being paid if 1
		if loanType == 1 {
			pmt = pmt / (1 + raterPerAnnum)
		}

	}

	return pmt
}
