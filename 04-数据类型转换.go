package main

import "fmt"

func main() {
	a := 5.0
	b := int(a)
	fmt.Printf("%T,%f\n", a, a) //float64,5.000000
	fmt.Printf("%T,%d\n", b, b) //float64,5.000000
}
