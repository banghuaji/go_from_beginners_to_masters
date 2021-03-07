package main

import (
	"fmt"
	"math"
)

//定义一个几何图形的接口
type geometry interface {
	area() float64       //面积
	perimeters() float64 //周长
}

type rect struct { //矩形
	length float64
	width  float64
}

type circle struct { //圆形
	radius float64
}

//方法的实现：
func (r rect) perimeters() float64 {
	return 2 * (r.length + r.width)
}

func (r rect) area() float64 {
	return r.width * r.length
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) perimeters() float64 {
	return math.Pi * c.radius * c.radius
}

//如何将接口作为参数传递
func PaaSInterfaces(g geometry) {
	fmt.Println("表明成功的实例化了USB的接口")
}
func main() {
	//查看接口如何调用
	r := rect{width: 3, length: 4}
	fmt.Println("周长是：", r.perimeters())
	fmt.Println("面积是：", r.area())

	a := circle{}
	a.radius = 7

	fmt.Println("周长是：", math.Trunc(a.perimeters()*1e2+0.5)*1e-2)
	fmt.Println("面积是：", fmt.Sprintf("%.2f", a.area()))
	fmt.Println(1.00 << 3)

	//定义一个接口
	var a1 geometry
	a1 = rect{width: 34, length: 243}
	fmt.Println("赋值给接口的面积是：", a1.area())
	PaaSInterfaces(a1)

	var a23 Usb
	a23 = PhoneConnect{"APPLE"}
	a23.Connected()
	DisConnected2(a23)
	DisConnected2(45)

}

type Usb interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connected()
}

type empty interface {
}

type PhoneConnect struct {
	name string
}

func (p PhoneConnect) Name() string {
	return p.name
}

func (p PhoneConnect) Connected() {
	fmt.Println("Connected:", p.name)
}
func DisConnected(usb Usb) {

	if v, ok := usb.(PhoneConnect); ok {
		fmt.Println("DisConnected:", v.name)
		return
	}
	fmt.Println("DisConnected: unknow")

}

//验证入参是一个支持所有类型的接口
func DisConnected2(usb interface{}) {

	switch v := usb.(type) {
	case int:
		fmt.Println("this is a int type ,value is :", v)
	case PhoneConnect:
		fmt.Println("DisConnected:", v.name)
	default:
		fmt.Println("DisConnected: unknow")
	}

}
