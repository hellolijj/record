package main

import (
	"fmt"
	"bufio"
)

type ByteCounter int

func (c *ByteCounter)Write(p []byte)(int, error)  {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func testByteCounter() {
	var c ByteCounter
	c.Write([]byte{'c', 'a'})
	c.Write([]byte{'c', 'a', 'b'})
	fmt.Fprintf(&c, "%s", "abcd")
	fmt.Println(c)
}

type WorldCounter int

func (c *WorldCounter)Write(p []byte)(int, error)  {
	length, _, _:= bufio.ScanWords(p, true)
	*c += WorldCounter(length)
	return len(p), nil
}

func testWorldCounter()  {
	var worldCounter WorldCounter
	worldCounter.Write([]byte("hello world adfa"))
	worldCounter.Write([]byte("hello world this is"))
	fmt.Println(worldCounter)
}

func main()  {
	testByteCounter()
	testWorldCounter()
}

