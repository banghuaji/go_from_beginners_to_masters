package main

import (
	"fmt"
)

func main() {
	SliceKeng()

}


//坑  slice,
//slice是地址传递，但是slice在地址空间不够厚会进行地址的重新new，
//使用中需要将值传递到其中
func SliceKeng() {
	s := make([]int, 0)
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)
}

func Pingpong(s []int) []int {
	return append(s, 2)
}
