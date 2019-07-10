package main

import "fmt"

/////////////////// 普通参数 ///////////////////
func test(a int, b int) {
	fmt.Println(a, b)
}

//////////////////// 不定参数 //////////////////////
func sum_func(args ...int) int {
	n := len(args)
	fmt.Println(n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += args[i]
	}
	fmt.Println(sum)
	return sum
}

//////////////////// 函数嵌套 //////////////////////
func enter(x int, arr ...int) {
	sum_func(arr[:]...) //因为sum_func没有定义返回值，计算一直不成功
}

/////////////////// 函数返回值 ////////////////////
func div(x int, arr ...int) int {
	s := sum_func(arr[:]...)
	divide := x / s
	return divide
}

func main() {
	test(3, 5)
	sum_func(1, 2, 3, 4)
	enter(6, 2, 3)

	result := div(6, 2, 3)
	fmt.Println("函数嵌套除法：", result)
}
