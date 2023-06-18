package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var c, python, java bool
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
	//swap
	a, b := swap("a", "b")
	fmt.Println(a, b)
	x, y := split(12)
	fmt.Println(x, y)
	var i int
	fmt.Println(i, c, python, java)

	var ii, j int = 1, 2
	var apple, babana = true, "no!"
	fmt.Println(ii, j, apple, babana)

	fmt.Println(":=  可在类型明确的地方替代var")
	fmt.Println("函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。")
	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // uint8 的别名

		rune // int32 的别名
		    // 表示一个 Unicode 码点

		float32 float64

		complex64 complex128
	*/
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("没有明确初始值的变量声明会被赋予它们的 零值。")
	var one int
	var two bool
	var three float64
	var four string
	fmt.Printf("%v %v %v %q\n", one, two, three, four)

	var iz int = 42
	var f float64 = float64(i)
	var ff float64 = 42
	var ui uint = 42
	fmt.Println(iz, f, ff, ui)

	//var ilx int
	/*j := ilx
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128*/
	const Pi = 3.14

}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 无参数的return 返回 已命名的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
