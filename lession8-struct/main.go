package main
//结构演示 struct
import "fmt"

func main() {
	a := test{}
	fmt.Println(a)
	a1 := persion{}
	//赋值初始化
	a1.name = "jone"
	a1.age = 12
	fmt.Println(a1)

	//字面值的初始化
	a2 := persion{name: "aaa", age: 34}
	fmt.Println(a2)

	//结构也是一个值类型,是值拷贝
	A(a2)
	fmt.Println(a2)

	//进行指针传递，修改会导致原先类修改
	P(&a2)
	fmt.Println(a2)

	//指针类型的初始化
	fmt.Println("指针类型的初始化")
	a3 := &persion{name: "aaa", age: 34}
	fmt.Println(a3)
	P(a3)
	a3.name = "icu"
	fmt.Println(a3)

	//匿名结构
	nimingjiegou()
	//匿名结构在结构体中
	fmt.Println("匿名结构在结构体中")
	superNimingHanShu()
}

func superNimingHanShu() {
	a:=student{}
	a.age = 17
	a.name = "xiaoli"
	a.contract.city = "shanghai"
	a.contract.phone = "110"
	a.persion.name = "100"
	a.persion.age = 100
	a.age =1000
	fmt.Println(a)
}

type student struct {
	name     string
	age      int
	contract struct {
		city  string
		phone string
	}
	persion
}

func nimingjiegou() {
	a := struct {
		name string
		age  int
	}{}
	fmt.Println(a)
}

type test struct {
}

type persion struct {
	name string
	age  int
}

func A(per persion) {
	per.age = 100
	fmt.Println(per)
}

func P(per *persion) {
	per.age = 101
	fmt.Println(per)
}
