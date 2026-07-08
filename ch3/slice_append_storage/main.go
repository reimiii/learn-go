package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	fmt.Println(x)

	y := x[:2]
	fmt.Println(y)

	fmt.Println(cap(x), cap(y))
	y = append(y, "z")

	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
