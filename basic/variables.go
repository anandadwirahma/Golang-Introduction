package basic

import "fmt"

func Variables() {
	var c, pyhton, java bool

	var i int
	fmt.Println(i, c, pyhton, java)
}

func VarInitializers() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func VarShortDeclare() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

//** NOTES **\\
// - "Variable Intializers", var declaration ca include intializers, one per variable
// - if an initializer is present, the type can be omitted. variable will take the type of the initalizer.
