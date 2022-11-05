package main

import "fmt"

func test() (int, int) {
	//定义一个test函数，它会返回两个int类型的值，每次调用将会返回100 和200两个数值。这里我们先不用管函数的定义,我们用这个函数来理解_这个匿名变量。
	return 100, 200
}

func main() {
	fmt.Println("hello world")
	//var 变量名字 变量类型=
	var name string = "jiangyiheng"
	//自动识别变量类型的语法糖
	age := 28
	salary := 33.3
	var count float64 = 3333.3
	fmt.Println(age, salary, count)
	//输出变量的类型 语法糖浮点默认使用float64
	fmt.Printf("%T,%T,%T", age, name, salary)
	fmt.Printf("内存地址:%p", &age) //取地址&变量名
	println()
	fmt.Println(salary, "-----------", count)
	//进行对值之间进行交换，如果类型不同需要转换
	salary, count = count, salary
	fmt.Println(salary, "-----------", count)

	a, _ := test()
	_, b := test()
	//在第一行代码中，我们只需要获取第一个返回值，所以第二个返回值定义为匿名变量_
	//在第二行代码中，我们只需要获取第二个返回值，所以第一个返回值定义为匿名变量―
	fmt.Println(a, b)

}
