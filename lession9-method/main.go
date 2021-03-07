package main

import "fmt"

//结构体A
type A struct {
	TZ
	B
	name string
	age  int
}

//结构体A的方法
func (a A) Print() {
	fmt.Println(a.name)
	//方法也是能够作为继承
	a.PlaneB()
}

//结构体A的方法,此处的A也是遵循值传递
func (a A) Print2() {
	fmt.Println(a.name)
	a.name = "xiaoji"
	fmt.Println(a)
}
//结构体A的方法,此处的A也是遵循值传递
func (a *A) Print3() {
	fmt.Println(a.name)
	a.name = "xiaoji"
	fmt.Println(a)
}


func main() {
	fmt.Println("结构体增加一个方法")
	a := A{name: "xiaoming"}
	a.Print()

	fmt.Println("结构体方法时一个值传递")
	a.Print2()
	fmt.Println(a)

	fmt.Println("结构体方法，传递一个指针")
	a.Print3()
	fmt.Println(a)





	tz := TZ(123)
	tz = 345
	tz.Print2()
}


type TZ int

func (a TZ) Print2() {
	fmt.Println(a)
}

func (a *TZ) Increase(num int) {
	*a = *a + TZ(num) //需要做一次强制转换
	//*a = 100
	//fmt.Println(a)
}

type B struct {
	city string
	six  int
}

func (b B)PlaneB()  {
	fmt.Println("this is a metheus B" )
}