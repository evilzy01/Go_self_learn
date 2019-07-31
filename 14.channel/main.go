package main

import (
	"fmt"
)

type Cat struct {
	name string
	age  int
}

func main() {
	////////////////  1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值 =%v \n", intChan)

	//注意，当给管道写入数据时，不能超过其容量
	intChan <- 10
	num := 211
	intChan <- num

	fmt.Println("channel length", len(intChan))
	fmt.Println("channel capacity", cap(intChan))

	//读出数据，注意读的时候管道有没有空，会deadlock
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 is", num2)

	fmt.Println("channel length", len(intChan))
	fmt.Println("channel capacity", cap(intChan))

	////////////////// 2. 创建一个可以存放10个map类型的管道
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)

	m1 := make(map[string]string, 20)
	m1["city1"] = "beijing"
	m1["city2"] = "shanghai"

	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"

	mapChan <- m1
	mapChan <- m2

	////////////////// 3. 创建一个可以存放任意类型的
	var mixChan chan interface{}
	mixChan = make(chan interface{}, 10)

	mixChan <- 10
	mixChan <- "hhh"

	cc := Cat{name: "kitty", age: 4}
	mixChan <- cc

	<-mixChan
	<-mixChan

	Newcat := <-mixChan
	//以下写法是错误的，编译不通过
	//	fmt.Println(Newcat.name)
	a := Newcat.(Cat) // 类型断言？？？？
	fmt.Println(a.name)

	///////////////// 遍历管道
	close(mapChan)
	for v := range mapChan {
		fmt.Println(v)
	}
}
