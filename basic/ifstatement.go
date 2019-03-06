package basic

import (
	"fmt"
	"math"
)

func Ifstatement(x float64) string {
	if x < 0 {
		return Ifstatement(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Ifshortstatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Ifelsestatement(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
