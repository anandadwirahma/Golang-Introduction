package basic

func FuncBasic(x int, y int) int {
	return x + y
}

func FuncContiued(x, y int) int {
	return x + y
}

func FuncMultipleResult(x, y string) (string, string) {
	return y, x
}

func FuncNameReturnValue(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//** NOTES **\\
// - "Funtion" can take zero or more aarguments.
//         -> Type in argument is type of input value arguments
//         -> Type in out parenthesis is type of value output funtion

// - "Funtion Continued" possible for writing type of argument just once if we need same data type in argument
// - You can omit the type from all but the last, like :
//        -> x, y int
