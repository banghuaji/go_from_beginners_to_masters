//常量与运算符
package main

import "fmt"

const a = 5
const b = a * 20
const (
	c, d = 7, 0
	MaxA = 1
	MaxB = iota
	MaxC = iota
)
const (
	c2, d2 = 7, 0
	MaxA2 = "A"
	MaxB2 = iota
	MaxC2 = iota
)

func main() {
	fmt.Println(MaxA)
	fmt.Println(MaxB)
	fmt.Println(MaxC)
	fmt.Println(MaxA2)
	fmt.Println(MaxB2)
	fmt.Println(MaxC2)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(5 % 7)

	aa := 2
	bb :=0
	if aa > 0 || (10/bb) > 1{
		fmt.Println("OK" )
	} else{
		fmt.Println("no ")
	}


}
