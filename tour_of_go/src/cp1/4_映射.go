package main

import (
	"fmt"
	"math"
)

type Vertexz struct {
	Lat, Long float64
}

var m map[string]Vertexz

func main() {
	m = make(map[string]Vertexz)
	m["Bell Labs"] = Vertexz{
		40.124123, -74.123123,
	}
	fmt.Println(m["bell Labs"])

	//map 的文法
	var mm = map[string]Vertexz{
		"A": {
			40.23423, 12.34523,
		}, "b": {
			123, 123,
		},
	}
	fmt.Println(mm)

	mi := make(map[string]int)
	fmt.Println(mi)
	mi["Answer"] = 42
	delete(mi, "Answer")
	v, ok := mi["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("函数也是值。它们可以像其它值一样传递。函数值可以用作函数的参数或返回值。")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("函数的闭包")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
