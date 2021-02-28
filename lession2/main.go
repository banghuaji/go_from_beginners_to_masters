//类型与变量

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var a int32
	var b float32
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var arr [8]bool
	fmt.Println(math.MinInt8)
	fmt.Println(arr)
	arr[4] = true
	fmt.Println(arr)

	//神明与赋值
	b3 := 122
	fmt.Println(b3)

	aa, bb, cc, dd := 1, "23", 3, 4.0000
	fmt.Println(aa, bb, cc, dd)

	var aaa float32 = 11.34343
	//var abool bool  =  false
	fmt.Println(int(aaa))
	//fmt.Println(string(abool))

	var ai11nt int = 65

	bstring := string(ai11nt)
	fmt.Println(bstring)
	fmt.Println(strconv.Itoa(ai11nt))
	fmt.Println(strconv.Atoi("42432432"))


}
