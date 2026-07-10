package main

import "fmt"

func main() {
	arrayConversion()
	arrayPointerConversion()
	panicArrayConversion()
}

func arrayConversion() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)

	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice, xArray, smallArray)
}

func arrayPointerConversion() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPoiner := (*[4]int)(xSlice)
	xSlice[0] = 10
	xArrayPoiner[1] = 20
	fmt.Println(xSlice, xArrayPoiner)
}

func panicArrayConversion() {
	xSlice := []int{1, 2, 3, 4}
	panicArray := [5]int(xSlice)
	fmt.Println(panicArray)
}
