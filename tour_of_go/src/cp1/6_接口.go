package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloatz(-math.Sqrt2)
	v := Vertexe{3, 4}
	a = f
	a = &v
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
	fmt.Println(a.Abs())

	fmt.Println("接口与隐式实现")
	/*
		类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。

		隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

		因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
	*/
	var i I = T{"Hello"}
	i.M()
}

type MyFloatz float64

func (f MyFloatz) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertexe struct {
	X, Y float64
}

func (v *Vertexe) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)

}
