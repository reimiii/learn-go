package main

import "fmt"

func main() {
	const value = 69
	var i int = value
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}
