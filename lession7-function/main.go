package main

import "fmt"

//slice是地址传递
//普通的类型是值传递，需要取出地址来操作。
//函数也是一种类型

func main() {
	a := clouser(10)
	fmt.Println(a(12))
	B(1, 2, 3, 4, 5, 6, 7)
	AAA("aaa",333)
}

func A(a string, b, c int) int {

	return 0
}
func AA(a string, b, c int) (int, int, int) {

	return 0, 0, 0
}

func AAA(a string, b ...int) (d int) {
	fmt.Println(a,b)
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