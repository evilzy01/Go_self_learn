package main

import (
	"fmt"
	"time"
)

//在主线程，开启一个goroutine，每隔一秒输出hello world
//在主线程中每隔一秒输出hello golang， 10次后退出程序
//要求协程和主线程同时执行

//一般协程是以函数为单位开启的
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world")
		time.Sleep(time.Second)
	}
}

func main() {

	//	test() //如果这样写，就是先打印完hello world再打印hello golang，即使他们是独立的两个过程，也在一个CPU上，低效率
	go test() // 开启一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("hello golang")
		time.Sleep(time.Second)
	}
}
