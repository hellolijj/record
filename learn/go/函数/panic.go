package main

import (
	"image/color"
	"time"
	"fmt"
	"math"
)

type Values map[string][]string

func (v Values)Get(key string) string  {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values)Add(key, val string) {
	v[key] = append(v[key], val)
}

func testValue()  {
	var v Values
	v = Values{
		"lijj": {"asfaf", "123"},
	}

	v.Add("lijj", "this is a big tu")
	//fmt.Println(v)

	v = nil
	fmt.Println(v.Get("lj"))
	//v.Add("af", "afaf")  // 报错

}

type Point struct {
	x, y int
}

func (p Point)Distance(q Point) float64 {
	return math.Sqrt(float64((p.y-q.y)*(p.y-q.y)+(p.x-q.x)*(p.x-q.x)))
}

func (p *Point)Scale(scale int)  {
	p.x *= scale
	p.y *= scale
}

func testPoint()  {
	var x, y Point
	x = Point{1, 2}
	y = Point{4, 5}

	fmt.Println(x.Distance(y))
}


type Path []Point

func (p Path)Distance() float64 {

	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func testPathDistance()  {

	p := Path{{0,0}, {1,1}, {2,2}, {3, 3}, {4, 4}}

	fmt.Println(p.Distance())
	
}

type ColorPoint struct {
	Point
	Color color.Color
	Time time.Duration
}

// 测试 结构体内嵌变量的 方法
func testColorPoint()  {
	p := ColorPoint{Point{1, 1}, color.Black, time.Second}
	q := ColorPoint{Point{4, 5}, color.White, 3 * time.Second}

	fmt.Println(p.Distance(q.Point))

	p.Scale(2)
	q.Scale(2)
	fmt.Println(p.Distance(q.Point))
}


func main()  {

	testValue()
	testPoint()
	testPathDistance()
	testColorPoint()



	//fmt.Println(v.Get("lijj"))

}