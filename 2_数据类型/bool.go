package main

import (
	"fmt"
	"unsafe"
)

func main() {
	////////////////////// bool 类型 ///////////////////////////
	var a bool
	fmt.Println(a)

	var b = true
	fmt.Println(b)

	c := true
	fmt.Println(c)

	v2 := (1 == 2)
	fmt.Println(v2)
	////////////////////// 整型 ///////////////////////////
	// uint 无符号整型，不可以接受负数，   int 有符号整型
	d := 9
	fmt.Println(unsafe.Sizeof(d))

	////////////////////// 浮点类型 ///////////////////////////
	// float32, float64(默认)——小数点后15位

	////////////////////// 字符类型 ///////////////////////////
	var x byte = 'h'      // '' 字符， “” 字符串
	fmt.Println(x)        // 输出是数字
	fmt.Printf("%c\n", x) //输出格式为一个字符
	fmt.Printf("%T ", x)
	fmt.Println('\n')

	////////////////////// 字符串类型 ///////////////////////////
	var str1 string = "abc"
	fmt.Println(str1)

	str2 := "hahahah"
	fmt.Println(str2)

	var m = len(str1)
	fmt.Println(m)

	var str3 = "加油" // 一个汉字占3个字符
	fmt.Println(len(str3))

	fmt.Println(str1 + " " + str2)

	////////////////////// 字符 与 字符串的区别 ///////////////////////////
	ch := 'a'
	str := "a"
	fmt.Println(ch, str)
	fmt.Println(str1[0], str1[2])

	////////////////////// 复数类型 ///////////////////////////
	t := 64 + 3i
	fmt.Println(t)
	fmt.Println(real(t), imag(t))

	////////////////////// 输入 ///////////////////////////
	//	var z float64
	//	//	fmt.Scanf("%d", &z
	//	fmt.Scan(&z)
	//	fmt.Println(z)

	////////////////////// 类型转换 ///////////////////////////
	k, w, y := 12, 13, 18
	sum := k + w + y
	fmt.Println(sum / 3)
	fmt.Println(float64(sum) / 3)

}
