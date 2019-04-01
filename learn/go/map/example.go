package main

import "fmt"

func main() {

	// map 的定义
	ages := make(map[string]int)
	ages["li"] = 1
	ages["lijj"] = 3
	fmt.Println(ages)

	agess := map[string]int{
		"lijj": 2,
		"hello": 3,
	}
	for k, v := range agess {
		fmt.Println(k, v)
	}



	// get map elements
	if name, ok := agess["lij"]; !ok {
		fmt.Println("error")
	} else {
		fmt.Println(name)
	}
}

// 判断 a b两个map类型的值 是否相等
func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || av != bv {
			return false
		}
	}
	return true
}