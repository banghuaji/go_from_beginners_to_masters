package main

import (
	"fmt"
	main2 "jibanghua.cn/go_from_beginners_to_masters/lession1/init_go_project"
	"time"
)


func main() {
	fmt.Println("hello jibanghua.cn")

	main2.BeginnersToMaster("7899")

	fmt.Println(time.Now().AddDate(0,0,1).Format(time.RFC850))

	fmt.Println(main2.Nowtime().Add(1))
}
