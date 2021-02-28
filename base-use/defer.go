package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		defer func() {
			if i == 2 {
				a := 0
				fmt.Println(i / a)
			}
		}()
	}
}
