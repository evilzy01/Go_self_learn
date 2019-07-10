package main

import (
	"fmt"
)

// 定义一个cat 结构体，将各个字段、属性，放入其中进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Score [3]int
}

func main() {
	var cat1 Cat //创建一个 结构体变量
	cat1.Name = "Kuly"
	cat1.Age = 3
	cat1.Color = "white"
	fmt.Println("Cat1 = ", cat1)

}
