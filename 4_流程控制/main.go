package main

import "fmt"

func main() {
	//	score := 700
	//	if score == 700 {
	//		fmt.Println("去上清华啦！")
	//	}

	if score := 700; score == 700 { // if 支持初始化语句，用;分割判断语句
		fmt.Println("去上清华啦~~")
	}

	s := 500
	if s < 300 {
		fmt.Println("sad!")
	} else if s < 600 {
		fmt.Println("ok..")
	} else {
		fmt.Println("yeah")
	}

	//////////////// switch /////////////////
	grade := 85
	switch grade {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	switch CPU := 2; {
	case CPU <= 3:
		fmt.Println("small")
	case CPU <= 5:
		fmt.Println("medium")
	default:
		fmt.Println("large")
	}

	//////////////// FOR /////////////////
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	for i, data := range str {
		fmt.Println(i, data)
	}

	////////////////  跳转语句 ///////////////////
	for i := 1; i < 10; i++ { // break
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for j := 1; j < 10; j++ { /// continue
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("aaaaa") /////  goto跳转
	goto END
	fmt.Println("bbbbb")
END:
	fmt.Println("ccccc")

}
