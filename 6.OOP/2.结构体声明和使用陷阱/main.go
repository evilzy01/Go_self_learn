package main

import (
	"fmt"
)

type prople struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针  结构体是这三种，零值都是nil，没有分配空间，如果需要使用这样的字段，需要先make，才能使用
	slice  []int             //切片
	map1   map[string]string //切片
}

func main() {
	var p1 prople
	fmt.Println(p1)

	if p1.slice == nil {
		fmt.Println("ok1")
	}

	//p1.slice[0] = 100 ///会报错，！使用slice前一定要先make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	fmt.Println(p1)

	p1.Name = "hahha"
	p1.Age = 22

	p2 := p1
	p2.Age = 24
	fmt.Println(p1)
	fmt.Println(p2)

}
