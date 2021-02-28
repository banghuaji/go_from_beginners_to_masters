package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	SliceKeng()
	timekeng()

}

func timekeng() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006:Mon Jan _2 15:04:05 "))
	fmt.Println(t.Unix()) //返回当前的时间戳
	StringToTime()
	UnixToString()

}

//时间字符串转时间(string  —>  Time)
func StringToTime() {

	formatTimeStr := "2017 - 04 - 11 13:33:37"
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC
	}
}

// 时间戳转时间字符串
func UnixToString()  {
	timeUnix:=time.Now().Unix() //已知的时间戳
	formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr)//打印结果：2017-04-11 13:30:39
	fmt.Println(reflect.TypeOf(formatTimeStr))
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
