package main

import (
	"flag"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

/*
模拟客户端
 */

var url = flag.String("url", "http://www.baidu.com", "fetct url")

func main()  {

	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch error: %v", err)
		os.Exit(1)
	}
	fmt.Println(*resp)
	//
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch body error: %v", err)
		os.Exit(2)
	}
	fmt.Println(string(b))
	fmt.Println(*url)
}