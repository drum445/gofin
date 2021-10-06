package gofin

import "math"

func pvF(rate, nper float64) float64 {
	return math.Pow((1 + rate), nper)
}

func pvAnnuity(rate, nper, pmt, fv, loanType float64) float64 {
	return (fv + pmt*(1+rate*loanType)*((math.Pow((1+rate), nper))-1)/rate)
}

func PV(rate, nper, pmt, fv, loanType float64) float64 {
	pv := 0.00
	if rate == 0.00 {
		pv = -pmt*nper - fv
	} else {
		raterPerAnnum := rate

		annuity := pvAnnuity(raterPerAnnum, nper, pmt, fv, loanType)
		fv := pvF(raterPerAnnum, nper)
		pv = -1 * (annuity / fv)
	}

	return pv
}

// PV*((1+ rate)^NPER)+ PMT*(1+rate*type)*(((1+ rate)^NPER)-1)/rate+FV = 0

// so rearranging to solve PV

//  PV = -1*(FV+ PMT*(1+rate*type)*(((1+ rate)^NPER)-1)/rate)/((1+ rate)^NPER)
