package main

import (
	"fmt"
)

func modify(m map[string]int) {
	m["a"] = 5
	fmt.Println(m)
}
func main() {
	m := map[string]int{}
	modify(m)
	fmt.Println(m)
	m["b"] = 65
	modify(m)
	fmt.Println(m)
	m["a"] = 10
	fmt.Println(m)
	modify(m)
	fmt.Println(m)
}
