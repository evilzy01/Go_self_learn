package main

import (
	"fmt"
)

//需求： 计算1-200各个数的阶乘，并放入到map中
//最后显示出来，要求用goroutine实现

//思路：
//1. 编写函数，计算阶乘，放入map中
//2. 启动的协程是多个，将结果放入map中，所以map是全局的

var intChan chan int

func test(n int) {
	intChan = make(chan int, 20)
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		intChan <- res
	}

}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i) // 启动200个协程
	}

	close(intChan)
	for i := range intChan {
		fmt.Println(i) //这样输出v会有很多0，因为阶乘超过了int可存范围
	}

}
