// 方法是作用在指定数据类型上的，    自定义类型（type）都可以有方法，不仅仅是struct

package main

import (
	"fmt"
)

type A struct {
	Num int
}

func (a A) test() { //表示A结构体有一个方法 test， （a A表示这个方法只能被A结构体类型调用）
	a.Num = 10
	fmt.Println(a.Num)
}

//// 给Person结构体添加一个speak方法，输出 ***是个好人
type Person struct {
	Name string
}

func (man Person) speak() {
	fmt.Println(man.Name, "is a good man")
}

//// 给Person结构体添加一个calculate方法, sum(1...100)
func (cal Person) calculate() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(cal.Name, "calcultate is ", sum)
}

//// 给Person结构体添加一个calculate方法, sum(1...n)
func (cal Person) calculate2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(cal.Name, "calcultate2 is ", sum)
}

//// 给Person结构体添加一个getsum方法, 可以计算两个数的和，并返回结果
func (cal Person) getsum(n int, m int) int {
	sum := m + n
	return sum
}

func main() {
	var p A
	p.test() //调用方法，只能通过p调用，并不是一个独立的函数
	fmt.Println(p.Num)

	var x Person
	x.Name = "jack"
	x.speak()

	x.calculate()
	x.calculate2(1000)

	s := x.getsum(2, 3)
	fmt.Println(s)
}
