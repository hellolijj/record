# 结构体

## 匿名变量

Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。

```go
type Point struct {
    X, Y int
}
type Circle struct {
    Point
    Radius int
}
type Wheel struct {
    Circle
    Spocks int
}
```

匿名变量的赋值不能通过如下方式
```go
var wheel Wheel
wheel.X = 3
```
而应该通过匿名变量特有的赋值方式
```go
w = Wheel{Cicle{Point{3, 4}, 5}, 6}
w = Wheel{
    Circle: Circle{
        Point: Point{
            X: 8, 
            Y: 8
        },
        Radius: 5,
    },
    Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
}
w.X = 43
```
- 匿名变量不能包含两个类型相同的类型成员变量  
- 因为 Cicle 是可导出的，可以使用匿名访问变量。如果 cicle 则不能使用简短的匿名访问变量
