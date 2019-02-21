package main

import (
	"fmt"
)
// ** Define Variable ** \\
// 1. Type 1 :
// var variableName type (define a variable with name "variableName" and "type")
// 2. Type 2 :
// var vname1, vname2, vname3 type (define three variables which type are "type")
// 3. Type 3 :
// var varaibleName type = values (define a variable with name "variableName", type "type" and value "value")
// 4. Type 4 :
// var vname1, vname2, vname3, type = v1, v2, v3
// (define three variables with type "type", and initialize their values.
// vname1 is v1, vname2 is v2, vname3 is v3)

// -- Define Variable with Short Assignment -- \\
// 1. Type 1 :
// vname1, vname2, vnam3 := v1,v2,v3
// (*notes :
// 	- only use inside of a function)
// 	- if you want to define global variable , use "var"
// 2. Type 2 :
// _, b := 34, 35 (_<blank> -> ignoring value to given in this parameter)

// ** Constant ** \\
// Variable which value cannot change during runtime
// * Example :
// const Pi = 3.1415926
// const i = 10000
// const MaxThread = 10
// const prefix = "astaxie_"

// ** Boolean **
// Value : true/false
// default values : false

func main() {
	fmt.Println("Test Boolean type : ")
	valid, available := booleanType()
	fmt.Println(valid)
	fmt.Println(available)

	fmt.Println("\nTest Integer type : ")
	int_value := integerType()
	fmt.Println(int_value)

	fmt.Println("\nTest String type : ")
	fmt.Println(stringType())

	fmt.Println("\nTest Change String : ")
	changeString()

	fmt.Println("\nTest iota enumerate")
	enumType()

	fmt.Println("\nTest Array Type : ")
	arrayType()

	fmt.Println("\nTest Array Two Dimension : ")
	arrayTwoDimension()

	fmt.Println("\nTest Slice : ")
	sliceType()
}

// ** Boolean Types ** \\

var isActive bool // global variable
var enabled, disabled = true, false

func booleanType() (bool, bool) {
	var available bool

	valid := false
	available = true

	return available, valid
}

// ** Integer Types ** \\

func integerType() int8 {
	var a int8
	var b int8

	a = 3
	b = 28
	c := a + b

	return c
}

// ** String Types ** \\

var frenchHello string
var emptyString string = ""

func stringType() (string,string,string,string,string) {
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiou"
	frenchHello := `Bonjour`

	return no, yes, maybe, japaneseHello, frenchHello
}

func changeString() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}

// ** Defined Variable By Group ** \\

func definebyGroup(){

	const (
		i = 100
		pi = 3.14515
		prefix = "Go_"
	)

	// var(
	// 	i int
	// 	pi float32
	// 	prefix string
	// )
}

// ** Iota Enumerate Types ** \\
// iota is keyword to make enum, it begins with 0, increased by 1

func enumType(){

	const(
		x = iota
		y = iota
		z = iota
		w // If there is no expression after the constants name, it uses the last expression,
		//so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
	)

	const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

	const (
		e, f, g = iota, iota, iota // values of iota are same in one line.
	)

	fmt.Println("values of x,y,z,w,v = " , x,y,z,w,v)
	fmt.Println("values of e,f,g = " , e,f,g)

}

func arrayType(){

	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	a := [3]int{1,2,3}
	b := [10]int{1,2,3}
	c := [...]int{4,5,6}

	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])
	fmt.Printf("A element is %d\n", a)
	fmt.Printf("B element is %d\n", b)
	fmt.Printf("C element is %d\n", c)
}

func arrayTwoDimension(){

	doublArray := [2][4]int{[4]int{1,2,3,4}, [4]int{5,6,7,8}}
	easyArray := [2][4]int{{1,2,3,4},{5,6,7,8}}

	fmt.Printf("Array two-Dimesional (hard practice) : %d\n", doublArray)
	fmt.Printf("Array two-Dimesional (easy practice) : %d\n", easyArray)
}

func sliceType() {

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a,b []byte
	a = ar[2:3]
	b = ar[3:5]

	fmt.Printf("Slice type A : %s\n", a)
	fmt.Printf("Slice type B : %s\n", b)
}