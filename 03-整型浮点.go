package main

import "fmt"

func main() {
	//定义一个整型
	var age int = 18
	fmt.Printf("%T,%d\n", age, age)

	//定义一个浮点型
	//默认是六位小数打印3.140000
	var money float32 = 3.19
	//保留小数---四舍五入 %.3  ------3.2
	fmt.Printf("%T,%.1f\n", money, money)
}
