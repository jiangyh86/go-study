package main

import "fmt"

// 全局变量
var name string = "蒋易恒"

//语法糖无法使用
//age:=23

func main() {
	name := "jyh"
	fmt.Println(name)

	//常量的定义,不可改变，const的值没有被使用不会报错
	const URL = "www.baidu.com"
	const URL2 = "www.google.com"
	const a, b, c = 12, 34, 45
	fmt.Println(URL)
	fmt.Println(URL2)
	fmt.Println(a, b, c)
}
