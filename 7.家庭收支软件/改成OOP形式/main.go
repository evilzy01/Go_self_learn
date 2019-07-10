// 把记账软件功能，封装到一个结构体中，然后通过调用该结构体方法，来实现记账、显示明细。
//结构体：FamilyAccount

package main

import (
	"fmt"
)

type FamilyAccount struct {
	//声明必要的字段

	//声明一个变量，来保存用户的输入选项
	key string
	//定义个变量，来控制是否退出for循环
	loop bool
	//定义账户余额
	balance float64
	//定义每次收支金额
	money float64
	//每次收支的说明
	notes string

	details string
}

//给该结构体绑定相应的方法
//显示主菜单
func (menu *FamilyAccount) mainmenu() {
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

		fmt.Scanln(&menu.key)

		switch {
		case menu.key == "1":
			menu.showDetails()
		case menu.key == "2":
			menu.income()
		case menu.key == "3":
			menu.pay()

		case menu.key == "4":
			menu.exit()

		default:
			fmt.Println("请输入正确选项")
		}
		if !menu.loop {
			break
		}
	}
}

//将显示明细写成一个方法
func (detail *FamilyAccount) showDetails() {
	fmt.Println("----------------收支明细记录...---------------")
	fmt.Println(detail.details)
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)

	this.balance += this.money

	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.notes)

	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.notes)
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("本次余额不足")
		//break
	}

	this.balance -= this.money

	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.notes)

	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.notes)
}

func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

func main() {
	var a FamilyAccount
	a.mainmenu()
}
