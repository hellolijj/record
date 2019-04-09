package main

import (
	"flag"
	"time"
	"fmt"
	"sort"
)

// flag int 用于使用flag的形式赋给一个int的值
var period = flag.Duration("period", 2*time.Second, "sleep period")
var n = flag.Int("n", 3, "n value")

func main()  {

	flag.Parse()
	fmt.Println("sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println(*n)
	fmt.Println()
	sort.Ints([]int{3,2,1,4,5,7});

}

/*
flag 的使用方法如下

定义
var n = flag.Int("n", 3, "n value")

程序使用
fmt.printf(*n)

交互使用
./flag -n 5
*/
