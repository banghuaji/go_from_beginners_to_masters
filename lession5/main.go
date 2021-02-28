//数组
package main

import "fmt"

func main() {

	//定义 var <varname> [n]type n >=0

	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b)

	//赋值
	bIntArray := [2]int{12, 32}
	fmt.Println(bIntArray)
	bIntArray2 := [...]int{12: 32}
	fmt.Println(bIntArray2)

	var a1 [2]int = [2]int{1, 1}
	var b1 [2]int = [2]int{2, 2}
	fmt.Println(a1)
	fmt.Println(b1)

	//二维数组
	var a11 [2][2]int = [2][2]int{{1, 1}, {2, 2}}
	fmt.Println(a11)

	//切片slice
	var s1 []int
	fmt.Println(s1) //空slice

	aaa := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(aaa)
	s11 := aaa[5:10]      //取得 a[5,6,7,8,9]
	s11 = aaa[5:]         //取得 a[5,6,7,8,9]
	s11 = aaa[5:len(aaa)] //取得 a[5,6,7,8,9]
	fmt.Println(s11)

	ss1 := make([]int, 3, 10)
	ss1 = append(ss1, 11)
	fmt.Println(ss1)

	aaaa := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	//sa := aaaa[3:5]
	//fmt.Println(sa)
	bbbb := aaaa[0:]
	fmt.Println(bbbb)
	fmt.Printf("%v,%p\n",aaaa,aaaa)
	fmt.Printf("%v,%p\n",bbbb,bbbb)
}
