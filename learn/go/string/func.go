package main

import (
	"fmt"
	"strings"
	"strconv"

	"enev/exam"

)

func main()  {

	fmt.Println(strings.Contains("abc", "a"))

	fmt.Println(strings.Index("abc", "i"))

	fmt.Println(strings.Count("ababab", ""))

	fmt.Println(strconv.Itoa(3))

	b := fmt.Sprintf("%d", 2344)     // the way of int to string
	fmt.Println(b)

	fmt.Println(strconv.Itoa(234))  // int to ascii
	strconv.ParseInt("234", 10, 64)

	fmt.Println(exam.Printlens())

	fmt.Println(strings.Map())


}

/*
fmt.Sprintf  // 返回给变量
fmt.Printf  // 普通输入到终端
fmt.Fprintf   // 输出给制定到设备
 */

