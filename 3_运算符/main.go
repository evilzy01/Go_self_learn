package main

import "fmt"

func main() {
	// 算术运算符
	a := 10
	fmt.Println(a)

	a++
	fmt.Println(a)

	a += 2
	fmt.Println(a)

	//关系运算符
	b := 20
	fmt.Println(a == b)

	//逻辑运算符
	fmt.Println(!true) // &&与 ||或

	//位运算符 & | ^ << >>
	fmt.Println(a & b)

	fmt.Println(&a) ////这里的&是取址符号
	fmt.Printf("%p\n", &a)

}
