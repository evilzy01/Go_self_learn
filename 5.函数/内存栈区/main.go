package main

import "fmt"

func add(s1 int, s2 int) {
	sum := s1 + s2
	fmt.Println(sum)
}

//定义函数类型，为已存在的函数起别名
type FUNCDEMO func(int, int)

func main() {
	a := 10
	b := 20
	add(a, b)
	fmt.Println(add) //函数名 —— 函数在代码区的地址
	x := 10
	fmt.Println(&x)
	fmt.Printf("%T\n", add) ///函数类型

	var f FUNCDEMO //为什么要重新定义？？？之后再讲
	f = add
	fmt.Println(f)

	/////////////////// 函数作用域 /////////////////////////

}
