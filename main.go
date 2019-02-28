package main

import (
	"fmt"
	"github/Golang-Introduction/basic"
)

func main() {
	fmt.Println("==== Chapter String ====\n")

	fmt.Println("- Import :")
	basic.ImportMain()

	fmt.Println("- Export :")
	basic.ExportFunc()

	fmt.Println("- Function Basic :")
	fmt.Println(basic.FuncBasic(1, 2))

	fmt.Println("- Function Continued :")
	fmt.Println(basic.FuncContiued(3, 4))

	fmt.Println("- Function Multiple Result :")
	fmt.Println(basic.FuncMultipleResult("Thinks", "Smart"))

	fmt.Println("- Variable :")
	basic.Variables()

	fmt.Println("- Variable Initializers :")
	basic.VarInitializers()

	fmt.Println("- Variable Initializers :")
	basic.VarShortDeclare()

	fmt.Println("- Type Basic :")
	basic.TypeBasic()

	fmt.Println("- Type With Zero Values :")
	basic.TypeZeroVal()

	fmt.Println("- Type Conversation :")
	basic.TypeConversions()

	fmt.Println("- Type Inference :")
	basic.TypeInference()

	fmt.Println("- Constant Variable :")
	basic.Constant()

	fmt.Println("- Constant Numeric Constant :")
	fmt.Println(basic.ConsNeedFLoat(basic.Small))

	fmt.Println("- Constant Numeric Constant :")
	fmt.Println(basic.ConsNeedFLoat(basic.Big))

	fmt.Println("- Constant Numeric Constant :")
	fmt.Println(basic.ConsNeedInt(basic.Small))

	// fmt.Println("- Method with pointer receiver : ")
	// v := method.Vertex{3, 4}
	// v.Scale(10)
	// fmt.Println(v.Abs())

	// fmt.Println("- Method with pointer receiver : ")

	// fmt.Println("Test Boolean type : ")
	// valid, available := basic.BooleanType()
	// fmt.Println(valid)
	// fmt.Println(available)

	// fmt.Println("\nTest Integer type : ")
	// int_value := basic.IntegerType()
	// fmt.Println(int_value)

	// fmt.Println("\nTest String type : ")
	// fmt.Println(basic.StringType())

	// fmt.Println("\nTest Change String : ")
	// basic.ChangeString()

	// fmt.Println("\nTest iota enumerate")
	// basic.EnumType()

	// fmt.Println("\nTest Array Type : ")
	// basic.ArrayType()

	// fmt.Println("\nTest Array Two Dimension : ")
	// basic.ArrayTwoDimension()

	// fmt.Println("\nTest Slice : ")
	// basic.SliceType()
}
