package init_go_project

import (
	"fmt"
	"os"
)

func BeginnersToMaster(str string) {
	fmt.Println(str)
	s := os.Args[0]
	fmt.Println(s)

	var (
		qq string = "abcd"
		//ee string = "abcd"
		//tt string = "abcd"
	)
	fmt.Println(qq)

}

const (
	PI  float32 = 3.14
	PI2 float32 = 3.14
)

var (
	a string = "abcd"
	b string = "abcd"
	d string = "abcd"
)

//一般类型申明
type nameType int

// 结构的申明
type student struct {
	name   string
	age    int
	scored float32
}

//接口的申明
type goland interface {
	getLastRunSult()
}
