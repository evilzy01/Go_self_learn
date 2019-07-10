package main

import (
	"fmt"
)

func main() {
	//声明一个变量，来保存用户的输入选项
	key := ""
	//定义个变量，来控制是否退出for循环
	loop := true
	//定义账户余额
	balance := 10000.0
	//定义每次收支金额
	money := 0.0
	//每次收支的说明
	notes := ""

	details := "收支\t账户金额\t收支金额\t说 明"

	for {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("------------------显示主菜单------------------")
		fmt.Println("                 1. 收支明细                  ")
		fmt.Println("                 2. 收入记录                  ")
		fmt.Println("-----------------3. 支出记录------------------")
		fmt.Println("-----------------4.  退出  ------------------")
		fmt.Println("请选择（1-4）：")

		fmt.Scanln(&key)

		switch {
		case key == "1":
			fmt.Println("----------------收支明细记录...---------------")
			fmt.Println(details)
		case key == "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)

			balance += money

			fmt.Println("本次收入说明：")
			fmt.Scanln(&notes)

			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, notes)

		case key == "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("本次余额不足")
				break
			}

			balance -= money

			fmt.Println("本次支出说明：")
			fmt.Scanln(&notes)

			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, notes)

		case key == "4":

			/////项目改进
			fmt.Println("你确定退出么？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("请重新输入y/n")

			}
			if choice == "y" {
				fmt.Println("-------------退出当前软件--------------")
				loop = false
			}

		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
}
