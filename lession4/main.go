//控制语句

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	var name *int = &a
	fmt.Println(name)

	if a1 := 1; a1 > 0 {
		//fmt.Println($i := a1)
	}

	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("=============================")
	aswitch := 0
	switch {
	case aswitch > 0:
		fmt.Println(0)
		fallthrough
	case aswitch > 1:
		fmt.Println(1)

	case aswitch > 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("NONE")
	}

	fmt.Println("=============BREAK语句================")
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("=============continue 语句================")

	for a := 1; a < 3; a++ {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i > 5 {
				continue
			}
		}
		fmt.Println("=============continue 语句 ================" + strconv.Itoa(a))
	}

	fmt.Println("=============数组 语句 ================")
	aint := [...]int{20:1,30:99}
	fmt.Println(aint)
}
