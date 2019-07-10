package main

import "fmt"

type abc struct {
	v int
}

func (a abc) aaaa() { //传入的是值，而不是引用
	a.v = 1
	fmt.Printf("1:%d\n", a.v)
}

func main() {

	//  指针的基本介绍
	var i int = 10
	fmt.Println("address of i:", &i)

	var ptr *int = &i
	fmt.Println("ptr is:", ptr)
	fmt.Println("address of ptr:", &ptr)

	fmt.Println("the value of pointer is:", *ptr)

	//指针案例和使用陷阱
	var num int = 5
	var ptr2 *int = &num
	*ptr2 = 11
	fmt.Println(num)

	var num2 int = 8
	ptr2 = &num2
	*ptr2 = 18
	fmt.Println(num2)

}
