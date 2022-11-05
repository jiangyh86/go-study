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
		k        //1
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}
