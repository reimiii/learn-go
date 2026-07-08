package main

import "fmt"

func main() {
	// var x [3]int
	// var xx = [3]int{10, 20, 30}

	// fmt.Println(x)
	// fmt.Println(xx)

	// var xxx = [12]int{1, 5: 4, 6, 10: 100, 15}
	// fmt.Println(xxx)

	// var yy = [...]int{1, 2, 3}
	// var y = [3]int{1, 2, 3}
	// fmt.Println(yy == y)

	// var dxa [2][3]int
	// fmt.Println(dxa)

	// y[0] = 88
	// fmt.Println(y)
	// fmt.Println(len(xxx), len(dxa[1]))

	// fmt.Println("slice")
	// var xs = []int{10, 20, 111}
	// var xxs = []int{1, 5: 4, 6, 10: 100, 15}
	// xs[0] = 99

	// fmt.Println(xs, xxs)

	// var sx = []int{11, 22, 33}
	// fmt.Println(sx)

	// var _default = []int{1, 5: 4, 6, 10: 100, 15}
	// fmt.Println(_default)

	// _default[0] = 132
	// fmt.Println(_default)

	// var maybe_true []int

	// fmt.Println(maybe_true == nil)
	//

	// x := []int{1, 2, 3, 4, 5}
	// y := []int{1, 2, 3, 4, 5}
	// z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}
	// fmt.Println(slices.Equal(x, y)) // prints true
	// fmt.Println(slices.Equal(x, z)) // prints false
	// fmt.Println(slices.Equal(x, s)) // does not compile
	//
	//

	// var x []int
	// x = append(x, 10, 2)
	// fmt.Println(x)

	// var x = []int{1, 2, 3}
	// x = append(x, 11, 22)
	// fmt.Println(x)

	// y := []int{40, 202}
	// x = append(x, y...)

	// fmt.Println(x)
	//

	// x := make([]int, 5)
	// x = append(x, 10)
	// fmt.Println(x)

	x := make([]int, 0, 10)
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)

	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))
}
