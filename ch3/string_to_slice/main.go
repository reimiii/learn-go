package main

import "fmt"

func main() {
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(string(bs[0]))
	fmt.Println(bs)
	fmt.Println(rs)
}
