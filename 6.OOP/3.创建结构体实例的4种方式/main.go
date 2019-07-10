package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式一

	//方式二
	p2 := Person{"Mary", 20}

	//方式三
	var p3 *Person = new(Person)
	(*p3).Name = "smith"
	(*p3).Age = 30

}
