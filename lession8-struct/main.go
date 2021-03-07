package main

//结构演示 struct
//结构也是一种类型，同种类型可以进行相关的赋值
import "fmt"

type persion1 struct {
	Name string
	age  int
}

func initStruct() {
	var persion11 persion1
	persion11.age = 11
	persion11.Name = "xiaoming"
	fmt.Println(persion11)
	persion12 := persion1{Name: "xiaoji", age: 18}
	fmt.Println(persion12)
}

func Propagation(p persion1) {
	p.age = 99
	fmt.Println(p)
}

func AddressPass(p *persion1) {
	p.age = 99
	fmt.Println(p)
}

type PersionNiming struct {
	Name string
	Age  int
	Contact struct{
		Phone string
		Address string
	}
}

func main() {
	fmt.Println("初始化结构体")
	initStruct()

	//结构也是一个值类型,是值拷贝
	fmt.Println("值传递验证")
	initPersion1 := persion1{Name: "xiaoji", age: 18}
	Propagation(initPersion1)
	fmt.Println(initPersion1)

	fmt.Println("地址传递验证")
	initPersionAddress := persion1{Name: "xiaoji", age: 18}
	AddressPass(&initPersionAddress)
	fmt.Println(initPersionAddress)

	fmt.Println("多次地址传递，可将类声明为结构指针类型，")
	PersionAddress := &persion1{Name: "xiaoji", age: 18}
	AddressPass(PersionAddress)
	fmt.Println(PersionAddress)

	fmt.Println("匿名函数展示")
	aPersion := struct {
		Name string
		age  int
	}{
		Name: "23323",
		age:  345,
	}
	fmt.Println(aPersion)
	niming := PersionNiming{Name: "xiaoji", Age: 18}
	niming.Contact.Phone = "110"
	niming.Contact.Address= "66666666"
	fmt.Println(niming)

	fmt.Println("======================================")

	//匿名结构
	nimingjiegou()
	//匿名结构在结构体中
	fmt.Println("匿名结构在结构体中")
	superNimingHanShu()
}

func superNimingHanShu() {
	a := student{}
	a.age = 17
	a.name = "xiaoli"
	a.contract.city = "shanghai"
	a.contract.phone = "110"
	a.person.name = "100"
	a.person.age = 100
	a.age = 1000
	fmt.Println(a)
}

type student struct {
	name     string
	age      int
	contract struct {
		city  string
		phone string
	}
	person
}

func nimingjiegou() {
	a := struct {
		name string
		age  int
	}{}
	fmt.Println(a)
}



type person struct {
	name string
	age  int
}

func A(per person) {
	per.age = 100
	fmt.Println(per)
}

func P(per *person) {
	per.age = 101
	fmt.Println(per)
}
