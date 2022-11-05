package main

import "fmt"

func main() {
	/*	//if语句
		a := true
		b := false
		if a {
			fmt.Printf("数据类型为:%T \n%v", b, b)
		}
		//if-else
		age := 44
		if age < 0 {

		} else if age > 50 {

		} else {
			fmt.Printf("\n%T %d\n", age, age)
		}

		//switch 语句
		switch age {
		case 42:
			fmt.Println("不是42")
		case 44:
			{
				fmt.Printf("年龄:%d", age)
			}
		}*/
	//Type Switch
	/*	var x interface{}

		switch i := x.(type) {
		case nil:
			fmt.Printf(" x 的类型 :%T", i)
		case int:
			fmt.Printf("x 是 int 型")
		case float64:
			fmt.Printf("x 是 float64 型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 型")
		case bool, string:
			fmt.Printf("x 是 bool 或 string 型")
		default:
			fmt.Printf("未知型")
		}
		age := 44
		//fallthrough,不能使用在 TypeSwitch 中
		switch age {
		case 47:
			if age == 44 {
				//break
			}
			fmt.Println("不是42")
			//fallthrough
		case 43, 45, 46:
			fmt.Printf("年龄:%d", age)
		default:
			fmt.Printf("默认输出语句")
		}*/
	var a, b, c int = 3, 4, 5
	fmt.Println(a, b, c)
}
