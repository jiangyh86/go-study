# 基础语法

## 一、Hello World

```go
package main

import "fmt"

func main() {
   fmt.Println("hello world")
}
```

## 二、定义变量

![image-20221105175348743](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105175348743.png)

![image-20221105175420624](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105175420624.png)

### 变量的可见性

- 声明在函数内部，是函数的本地值，类似private
- 声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
- 声明在函数外部且首字母大写是所有包可见的全局值,类似public

```go
package main

import "fmt"

func main() {
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

}

```

## 三、匿名变量

匿名变量的特点是一个下画线"_"，""本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它)，==但任何赋给这个标识符的值都将被抛弃==，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。例如:

```go
func test() (int, int) {
	//定义一个test函数，它会返回两个int类型的值，每次调用将会返回100 和200两个数值。这里我们先不用管函数的定义,我们用这个函数来理解_这个匿名变量。
	return 100, 200
}	
func main() {
	a, _ := test()
	_, b := test()
	//在第一行代码中，我们只需要获取第一个返回值，所以第二个返回值定义为匿名变量_
	//在第二行代码中，我们只需要获取第二个返回值，所以第一个返回值定义为匿名变量―
	fmt.Println(a, b)
}
```

## 四、局部全局变量、常量

```go
package main

import "fmt"

//全局变量
var name string = "蒋易恒"
//语法糖无法使用
//age:=23

func main() {
	name := "jyh"
	fmt.Println(name)
}
```

### 常量定义

```go
	//常量的定义,不可改变，const的值没有被使用不会报错
	const URL = "www.baidu.com"
	const URL2 = "www.google.com"
	const a,b,c = 12,34,45
	fmt.Println(URL)
	fmt.Println(URL2)
	fmt.Println(a,b,c)
```

### iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。iota是go语言的常量计数器
iota在const关键字出现时将被重置为0(const 内部的第一行之前),const 中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。

```go
package main

import "fmt"

func main() {
   //iota的认识
   const (
      a = iota  //0
      b         //1
      c         //2
      d = "jyh" //jyh
      e         //jyh
      f = 100   //100
      g         //100
      h = iota  //7
      i         //8
   )
   const (
      j = iota //0
      k  //1
   )
   fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}
```

### int类型

![image-20221105175712925](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105175712925.png)

### float

![image-20221105175750370](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105175750370.png)

```go
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
```

![image-20221105181051891](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105181051891.png)

### Go语言的格式化输出中%d%T%v%b等的含义

![image-20221105193538279](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105193538279.png)

## 五、数据类型转换

在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于Go语言不存在隐式类型转换，==因此所有的类型转换都必须显式的声明:==

```go
package main

import "fmt"

func main() {
	a := 5.0
	b := int(a)
	fmt.Printf("%T,%f\n", a, a) //float64,5.000000
	fmt.Printf("%T,%d\n", b, b)//float64,5.000000
}

```

类型转换只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型(将int16转换为int32)。当从一个取值范围较大的类型转换到取值范围较小的类型时(将int32转换为int16或将float32转换为int)，会发生精度丢失（截断)的情况。

## 六、运算符

### 算术运算符

![image-20221105182705629](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105182705629.png)

### 关系运算符

![image-20221105183125602](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105183125602.png)

### 逻辑运算符

![image-20221105183149203](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105183149203.png)

### 位运算符

![image-20221105183534687](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105183534687.png)

### 赋值运算符

![image-20221105184659787](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105184659787.png)

## 七、获取键盘输入

```go
package main

import "fmt"

func main() {
   //获取键盘输入的信息
   var name string
   var age int8
   //fmt.Scanf("%s %d", &name, &age)
    //阻塞等待键盘输入
   fmt.Scan(&name, &age)
   fmt.Println(name, age)
}
```

## 八、流程控制

### if控制

```go
package main

import "fmt"

func main() {
   //if语句
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
		fmt.Printf("\n%T %d", age, age)
	}
}
```

### switch 语句--==支持多条件匹配==

```go
//switch 语句
switch age {
case 42:
   fmt.Println("不是42")
case 44:
   {
      fmt.Printf("年龄:%d", age)
   }
}
//多条件匹配
case 43,45,46:
		fmt.Printf("年龄:%d", age)
```

可以使用break语句跳出

default默认执行的语句，==在所有case都不匹配后==一定最后执行，

```go
switch age {
case 44:
   if age == 44 {
      break
   }
   fmt.Println("不是42")
   fallthrough
case 43, 45, 46:
   fmt.Printf("年龄:%d", age)
    	default:
		fmt.Printf("默认输出语句")
}
```

#### Type Switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

```go
//Type Switch
var x interface{}

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
```

#### fallthrough

使用 fallthrough 会强制执行后面的 case 语句，==fallthrough 不会判断下一条 case 的表达式结果是否为 true。==

```go
age := 44
//fallthrough,不能使用在 TypeSwitch 中
switch age {
case 44:
   fmt.Println("不是42")
   fallthrough
case 43:
   fmt.Printf("年龄:%d", age)
}
```

### select语句

跟 switch-case 相比，select-case 用法比较单一，它仅能用于 信道/通道 的相关操作。==后续学习==

```txt
选择 {
    案例通信条款  ：
       语句（ s ）；      
    案例交流子句  ：
       statement ( s )；
    /* 你可以定义任意数量的大小写 */
    default : /* 任选 */
       statement ( s );
}
```

- 案例都必须是一个通信者

- 所有频道的表达都会被求值

- 所有被发送的邮件都被要求获得价值

- 如果任意通信示例可以，它就执行，其他被加载。

- 如果有多个案例都可以运行，选择会经常公平地执行。否则可以执行。其他不会执行

  ：

  1. 如果有默认子句，则执行该语句。
  2. 如果没有默认值，选择将否决或直接执行子句，可以运行；不会重新对频道求值。

## 九、循环语句

这是 for 循环的基本模型。

```go
for [condition |  ( init; condition; increment ) | Range]
{
   statement(s);
}
```

可以看到 for 后面，可以接三种类型的表达式。

1. 接一个条件表达式
2. 接三个表达式
3. 接一个 range 表达式

###  接一个条件表达式

==go中没有while==,取而代之的是

```go
a := 5
for a <= 7 {
   a++
   fmt.Println(a)
}
```

### 接三个表达式

```go
for a := 4; a <= 10; a++ {
   fmt.Println(a)
}
```

### 不接表达式：无限循环

在 Go 语言中，没有 while 循环，如果要实现无限循环，也完全可以 for 来实现。

当你不加任何的判断条件时， 就相当于你每次的判断都为 true，程序就会一直处于运行状态，但是一般我们并不会让程序处于死循环，在满足一定的条件下，可以使用关键字 `break` 退出循环体，也可以使用 `continue` 直接跳到下一循环。

下面两种写法都是无限循环的写法。

```go
for {
    代码块
}

// 等价于
for ;; {
    代码块
}
---------------------------------------
import "fmt"

func main() {
    var i int = 1
    for {
        if i > 5 {
            break
        }
        fmt.Printf("hello, %d\n", i)
        i++
    }
}
```

### 接 for-range 语句

遍历一个可迭代对象，是一个很常用的操作。在 Go 可以使用 for-range 的方式来实现。

range 后可接数组、切片，字符串等

由于 range 会返回两个值：索引和数据，若你后面的代码用不到索引，需要使用 `_`匿名变量 表示 。

```go
arr := [...]string{"world", "python", "go"}
for a, item := range arr {
   fmt.Printf("下标：%d hello, %s\n", a, item)
}
```

## 十、go to 语句进行跳转使用

goto 后接一个标签，这个标签的意义是告诉 Go程序下一步要执行哪里的代码。

```go
//go to 的使用方法
//go to 实现循环的效果
a := 0
flag:
if a < 5 {
   a++
   println(a)
   goto flag
}
```

```go
//使用 goto 实现 类型 continue的效果，打印 1到10 的所有偶数。
for i := 1; i < 11; i++ {
   if i%2 != 0 {
      goto flag
   }
   fmt.Println(i)
flag:
}
```

### 注意事项

goto语句与标签之间不能有变量声明，否则编译错误。

![image-20221105204719880](https://jyhload.oss-cn-shanghai.aliyuncs.com/img/image-20221105204719880.png)
