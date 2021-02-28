package main

import "fmt"

func main() {
	var m map[int]string = make(map[int]string)
	a := m[10]
	fmt.Println(a)




	//课堂作业
	m1:=map[int]string{1:"a",2:"b",3:"c"}
	m2 := make(map[string]int)
	for k,v :=range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
