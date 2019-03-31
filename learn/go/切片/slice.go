package main

import "fmt"

func main() {

	sl := []int{1,2,3,4,5}
	//append(sl, 3)
	fmt.Printf("%V ", sl)
	for _, v := range sl{
		fmt.Println(v)
	}

}