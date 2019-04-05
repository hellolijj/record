package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {

	s := "hello world"
	//fmt.Printf("%T", s)

	fmt.Println(s + "I am lijj")

	ss := "hello, 世界"

	fmt.Println(len(ss))


	fmt.Println(utf8.DecodeRuneInString(ss))
	fmt.Println(utf8.RuneCountInString(ss))  // 输出 9

	for i, r := range ss {                  // 循环9次
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// 强制转换
	fmt.Println(string(0x4eac))
	
}

/*
s[0] = 's'  // s[0] 不允许再修改
utf8.DecodeRuneInString()  rune size rune utf8编码所占的字节
  */

