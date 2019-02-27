package main

import (
	"fmt"
	"github/Golang-Introduction/basic"
	"github/Golang-Introduction/method"
)

func main() {
	fmt.Println("- Method with pointer receiver : ")
	v := method.Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	fmt.Println("- Method with pointer receiver : ")

	fmt.Println("Test Boolean type : ")
	valid, available := basic.BooleanType()
	fmt.Println(valid)
	fmt.Println(available)

	fmt.Println("\nTest Integer type : ")
	int_value := basic.IntegerType()
	fmt.Println(int_value)

	fmt.Println("\nTest String type : ")
	fmt.Println(basic.StringType())

	fmt.Println("\nTest Change String : ")
	basic.ChangeString()

	fmt.Println("\nTest iota enumerate")
	basic.EnumType()

	fmt.Println("\nTest Array Type : ")
	basic.ArrayType()

	fmt.Println("\nTest Array Two Dimension : ")
	basic.ArrayTwoDimension()

	fmt.Println("\nTest Slice : ")
	basic.SliceType()
}
