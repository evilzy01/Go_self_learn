// defer的价值在于，当函数执行完毕后（创建了资源——打开文件、获取数据库链接、锁资源），可以及时的释放那个创建的资源

package main

import "fmt"

func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)

	n1++
	n2++
	sum := n1 + n2
	fmt.Println("ok3 sum is", sum)
	return sum
}

func main() {
	//当执行到defer时候，暂时先不执行，会将defer后面的语句压入到独立的栈中
	//当函数执行完毕时候，再从defer栈中，按照先入后出的方式出栈，执行

	//defer把语句压入栈的同时，也会将相关的值压入其中，n1n2早就入栈，后边语句不产生影响
	res := sum(10, 20)
	fmt.Println("res is", res)
}
