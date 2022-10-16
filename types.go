package main

import "fmt"

func types() {
	const a int32 = 12
	const b float32 = 20.5
	var c complex128 = 1 + 4i
	var d uint16 = 14

	n := 42
	pi := 3.14
	x, y := true, false
	z := "go is questianable"

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
