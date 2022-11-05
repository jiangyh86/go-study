package main

import "fmt"

func main() {
	//获取键盘输入的信息
	var name string
	var age int8
	//fmt.Scanf("%s %d", &name, &age)
	//阻塞等待键盘输入
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
}
