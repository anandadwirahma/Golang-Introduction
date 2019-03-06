package basic

import "fmt"

func Defer() {
	defer fmt.Println("1.world")
	fmt.Println("Hello")
}

//** NOTES **\\
// - "Defer" will execute until the surrounding function returns,
