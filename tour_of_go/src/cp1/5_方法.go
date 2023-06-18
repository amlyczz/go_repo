package main

import (
	"fmt"
	"math"
)

type Vertee struct {
	X, Y float64
}

// 方法就是一类带特殊的 接收者 参数的函数。
func (v Vertee) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	fmt.Println("结构体方法")
	fmt.Println("值接受者")
	v := Vertee{3, 4}
	fmt.Println(v.Abs())
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abss())

	fmt.Println("指针接收者")
	vv := Vertee{3, 4}
	vv.Scale(10)
	fmt.Println(v.Abs())

	fmt.Println("结构体方法重写为函数")
	fv := Vertee{3, 4}
	Scale(&fv, 10)
	fmt.Println(fv)

}

type MyFloat float64

func (f MyFloat) Abss() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)

}
func (v *Vertee) Scale(f float64) {
	v.X = v.X * f
	v.Y *= f
}

func Abss(v Vertee) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertee, f float64) {
	v.X = v.X * f
	v.Y *= f
}