//map也是一种数据结构，key-value类型
package main

import "fmt"

func main() {
	//map 的声明，但是没有分配数据空间
	var a map[int]string
	//在使用map前，需要先make，给map分配数据空间
	a = make(map[int]string, 10) //10的意思是分配10个这样的空间，可以放10对int string
	a[1] = "kkkk"                // key不可以重复
	a[2] = "ggggg"
	fmt.Println(a) //map里的顺序既不是插入顺序，也不是数字顺序

	//////////////// 第二种方式 /////////////////////
	cities := make(map[string]string)
	cities["no2"] = "shanghai"
	cities["no1"] = "beijing"
	cities["no3"] = "nanjing"

	fmt.Println(cities)

	////////////////// 第三种方式 //////////////////
	hero := map[string]string{
		"hero1": "tony",
		"hero2": "spider",
	}
	fmt.Println(hero)

	////////////////// 案例 /////////////
	//存学生信息，对应每个学生学号，name,gender两个信息
	student := make(map[string]map[string]string)

	student["no1"] = make(map[string]string)
	student["no1"]["name"] = "tom"
	student["no1"]["gender"] = "boy"
	student["no1"]["address"] = "shanghai"

	student["no2"] = make(map[string]string)
	student["no2"]["name"] = "zoey"

	fmt.Println(student)

	//////////////// CRUD /////////////////////
	// C,U很好理解
	//删除key
	delete(cities, "no3")
	fmt.Println(cities)
	//查找
	val, findres := cities["no2"]
	fmt.Println(val)
	fmt.Println(findres)

	//////////// map的遍历 //////////////////////

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	for m, n := range student {
		for x, y := range n {
			fmt.Println(m, x, y)
		}
	}

	fmt.Println(len(cities))

}
