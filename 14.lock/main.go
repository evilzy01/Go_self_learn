package main

import (
	"fmt"
	"sync"
	"time"
)

//需求： 计算1-200各个数的阶乘，并放入到map中
//最后显示出来，要求用goroutine实现

//思路：
//1. 编写函数，计算阶乘，放入map中
//2. 启动的协程是多个，将结果放入map中，所以map是全局的

var (
	mymap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()    //所以要加锁
	mymap[n] = res //并发写，有问题
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i) // 启动200个协程
	}

	time.Sleep(time.Second * 10) //主线程等多久，是个问题？？

	lock.Lock()
	for i, v := range mymap {
		fmt.Println(i, v) //这样输出v会有很多0，因为阶乘超过了int可存范围
	}
	lock.Unlock()

}
