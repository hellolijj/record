package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "my name is lijj"
	a[1] = "this is a word"
	fmt.Println("hello", a)
	fmt.Printf("%T %v", a[0])


	b := [3]int{1,2,3}
	for i, v := range b {
		fmt.Printf("%d %d\n", i, v)
	}

}