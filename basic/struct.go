package basic

import "fmt"

type Vertex struct {
	X int
	Y int
}

func PointertoStructs() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

//** NOTES **\\
// - "Struct" is a collection of fields
