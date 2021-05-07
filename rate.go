package gofin

import (
	"fmt"
	"math"
	"strconv"
)

func Rate(nper, pmt, pv, fv, loanType float64) float64 {
	rate, _ := RateWithLimit(nper, pmt, pv, fv, loanType, 0)
	return rate
}

func RateWithLimit(nper, pmt, pv, fv, loanType, maxIterations float64) (float64, error) {
	guess := 0.1

	tolerancy := 1e-6
	isClose := false
	nextGuess := 0.00

	var iterations float64

	for !isClose {
		if maxIterations > 0 && iterations > maxIterations {
			return 0, fmt.Errorf("iterations limit reached")
		}
		temp := newtonIter(guess, nper, pmt, pv, fv, loanType)
		nextGuess, _ = strconv.ParseFloat(fmt.Sprintf("%.20f", guess-temp), 64)
		diff := math.Abs(nextGuess - guess)
		isClose = diff < tolerancy
		guess = nextGuess
		iterations += 1
	}

	return nextGuess * 12 * 100, nil
}

func newtonIter(r, n, p, x, y, w float64) float64 {
	t1 := math.Pow(r + 1, n)
	t2 := math.Pow(r + 1, n - 1)
	return (y + t1*x + p*(t1-1)*(r*w+1)/r) / (n*t2*x - p*(t1-1)*(r*w+1)/(math.Pow(r, 2)) + n*p*t2*(r*w+1)/r + p*(t1-1)*w/r)
}
