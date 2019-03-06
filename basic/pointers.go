package basic

import "fmt"

var numberA int = 4
var numberB *int = &numberA

func Pointer() {
	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address)   :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address)   :", numberB)
}

func PointerParam(original *int, value int) {
	*original = value
}

//** NOTES **\\
// - "Pointer" refference or address memory.
// - For ex : a variable of type integer has a value 4, so the meaning of pointer is address memory where value 4 is stored.
// 			  not the vlaue of 4 itself.

// "&" (ampersand) ==> The usual variable can also taken the pointer value.
//    			   ==> This method is called "Refferencing".
// "*" (asterisk)  ==> The original value of the pointer variable can also be retrieved.
//                 ==> This method is called "Dereferencing"
