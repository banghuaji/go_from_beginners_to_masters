package main

import (
	"fmt"
	main2 "jibanghua.cn/go_from_beginners_to_masters/lession1/init_go_project"
	"strconv"
	"strings"
	"time"
)

const NAME, AGE, DD = "jibanghua", 113, 1.123456789012345678901234567890
const SIX = -1

var big int64  = 100000000000
var small int8  = 127
var small2 int  = 256*256*256

var a, b, c = "a1", "b1", "c1"

func main() {

	fmt.Println(a)
	fmt.Println(a.StringVar(a))
	fmt.Println(small)
	fmt.Println(small2)

	fmt.Println(big)

	fmt.Println(strings.Contains("ababa", "aeb"))

	fmt.Println("hello jibanghua.cn")

	fmt.Println(NAME + "," + strconv.Itoa(SIX) + "," + strconv.Itoa(AGE) + ","+ strconv.FormatFloat(DD, 'E', -1, 32)	)

	fmt.Println(main2.PI)

	main2.BeginnersToMaster("7899")

	fmt.Println(time.Now().AddDate(0, 0, 1).Format(time.RFC850))

	fmt.Println(main2.Nowtime().Add(1))
}
