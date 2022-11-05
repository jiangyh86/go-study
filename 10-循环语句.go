package main

import "fmt"

func main() {
	/*	a := 5
		for a <= 7 {
			a++
			fmt.Println(a)
		}
		for a := 4; a <= 10; a++ {
			fmt.Println(a)
		}*/

	//arr := [...]string{"world", "python", "go"}
	//for a, item := range arr {
	//	fmt.Printf("下标：%d hello, %s\n", a, item)
	//}
	//go to 的使用方法
	//go to 实现循环的效果
	/*	a := 0
		flag:
			if a < 5 {
				a++
				println(a)
				goto flag
			}*/
	//使用 goto 实现 类型 continue的效果，打印 1到10 的所有偶数。
	for i := 1; i < 11; i++ {
		if i%2 != 0 {
			goto flag
		}
		//goto语句与标签之间不能有变量声明，否则编译错误。
		//var gg string = "ccc"
		fmt.Println(i)
	flag:
	}
}
