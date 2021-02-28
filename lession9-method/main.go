package main

import "fmt"

func main() {
	a := A{name: "xiaoming"}
	a.Print()
	//fmt.Println(a.Print())

	tz := TZ(123)
	tz = 345
	tz.Print2()
}

type A struct {
	name string
	int
}

//结构体A的方法
func (a A) Print() {
	fmt.Println(a.name)
}

//结构体A的方法,此处的A也是遵循值传递
func (a A) Print2() {
	fmt.Println(a.name)
}

type TZ int

func (a TZ) Print2() {
	fmt.Println(a)
}

func (a *TZ) Increase(num int) {
	*a = *a + TZ(num)
	//*a = 100
	//fmt.Println(a)
}
