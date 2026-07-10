package main

import (
	"fmt"
	"maps"
)

func main() {
	totalWins := map[string]int{}

	totalWins["sammy"] = 1
	totalWins["shark"] = 20
	totalWins["cats"] = 1
	fmt.Println(totalWins)

	totalWins["cats"]++
	fmt.Println(totalWins)

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	z := map[string]int{
		"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5,
	}
	fmt.Println(maps.Equal(z, n))
}
