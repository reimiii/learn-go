package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	fmt.Println("before:", x)

	y := x[:2]
	fmt.Println("before:", y)

	z := x[1:]
	fmt.Println("before:", z)

	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
