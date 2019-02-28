package basic

import "fmt"

const (
	Pi    = 3.14
	Big   = 1 << 100
	Small = Big >> 99
)

func Constant() {
	fmt.Println("Pi = ", Pi)
}

func ConsNeedInt(x int) int {
	return x*10 + 1
}

func ConsNeedFLoat(x float64) float64 {
	return x * 0.1
}

//** NOTES **\\
// - "Contant", are declaraed like variables.
// - constant can be character,string,boolean,or numeric
// - constant cannot be delcared using := syntax
