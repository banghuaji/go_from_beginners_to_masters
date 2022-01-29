package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	main2 "go_from_beginners_to_masters7777/lession1/init_go_project"
)

func main() {
	fmt.Println("hello jibanghua.cn")
	xlFile, err := xlsx.OpenFile("")
	fmt.Print(xlFile, err)

	main2.BeginnersToMaster("7899")
}
