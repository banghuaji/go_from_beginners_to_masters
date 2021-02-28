package main

import "fmt"

func main() {
	a := clouser(10)
	fmt.Println(a(12))
	B(1,2,3,4,5,6,7)
}

func A(a string, b, c int) int {

	return 0
}

//不定长变参
func B(a ...int) {
	fmt.Println(a)
}

//匿名函数

func C(a int) {
	are := func() {
		fmt.Println(a)
	}
	are()
}

//闭包
func clouser(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
