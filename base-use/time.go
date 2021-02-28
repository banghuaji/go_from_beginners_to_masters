package main

/**
golang获取当前时间、时间戳和时间字符串及它们之间的相互转换
*/
import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//timeKeng()
	deferUse()
}

func deferUse() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
			if i == 2 {
				a := 0
				fmt.Println(i / a)
			}
		}(i)
	}
}

func timeKeng() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006:Mon Jan _2 15:04:05 "))
	fmt.Println(t.Unix()) //返回当前的时间戳
	StringToTime()
	UnixToString()
	fmt.Println(time.Now().Day())

}

//时间字符串转时间(string  —>  Time)
func StringToTime() {
	formatTimeStr := "2017-04-11 13:33:37"
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		fmt.Println(formatTime) //打印结果：2017-04-11 13:33:37 +0000 UTC
	}
}

// 时间戳转时间字符串
func UnixToString() {
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39
	fmt.Println(reflect.TypeOf(formatTimeStr))
}
