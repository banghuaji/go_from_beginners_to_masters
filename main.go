package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	main2 "go_from_beginner33333333s_to_masters/lession1/init_go_project"
)

func main() {
	fmt.Println("hello jibanghua.cn")
	xlFile, err := xlsx.OpenFile("")
	fmt.Print(xlFile, err)

	main2.BeginnersToMaster("7899")
}
