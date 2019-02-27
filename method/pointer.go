package method

type Vertex struct {
	A, B int64
}

func (v Vertex) Abs() int64 { // -> method with receiver
	return v.A*v.B + v.A*v.B
}

//-- Method with pointer receiver --\\
// - Value of variable a & b modified with value of variabel f
// - pointer can modified value to which the receiver points

func (v *Vertex) Scale(f int64) {
	v.A = v.A * f
	v.B = v.B * f
}
