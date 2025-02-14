package main

import (
	"fmt"
)

type Walker interface {
	Walk()
}

type Dog struct {
	name string
	age  int
}

func (d Dog) Walk() {
	fmt.Printf("This dog name is %s and age is %d", d.name, d.age)
}

type Cat struct {
	name string
	age  int
}

func (c Cat) Walk() {
	fmt.Printf("This cat name is %s and age is %d", c.name, c.age)
}
func MakeWalk(w Walker) {
	w.Walk()
}
func PrintType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Others Type")
	}
}
func main() {
	dog := Dog{name: "Rex", age: 2}
	MakeWalk(dog)
	cat := Cat{name: "Tix", age: 3}
	MakeWalk(cat)
	PrintType(5)
	PrintType("Hello")
	PrintType(true)
	PrintType(1.3)
}
