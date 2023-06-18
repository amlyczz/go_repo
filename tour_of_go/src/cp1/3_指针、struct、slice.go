package main

import (
	"math/rand"
)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("指针")
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)

	*p = *p + 1
	fmt.Println(*p)
	fmt.Println("struct")

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	sp := &v
	sp.X = 1e9
	fmt.Println(v)

	fmt.Println("结构体文法")
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		pp = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, *pp)

	fmt.Println("数组")
	var a [10]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	fmt.Println("slice")
	s := primes[1:4]
	fmt.Println(s)

	fmt.Println("切片文法类似于没有长度的数组文法。")
	qq := []int{2, 3, 5, 7, 9, 11}
	fmt.Println(qq)
	r := []bool{true, false, false, true}
	fmt.Println(r)
	sb := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
	}
	fmt.Println(sb)

	ns := []int{2, 3, 5, 7, 11, 13}
	printSlice(ns)
	ns = ns[:4]
	fmt.Println(ns)
	//fmt.Println(ns[8:])

	var nils []int
	fmt.Println(nils, len(nils), cap(nils))
	if nils == nil {

		fmt.Println("nil")
	}

	fmt.Println("切片可以用内建函数 make 来创建，这也是你创建动态数组的方式")
	ma := make([]int, 5)
	printSlice2("ma", ma)
	mb := make([]int, 0, 5)
	printSlice2("b", mb)

	mc := mb[:2]
	printSlice2("mc", mc)
	md := mc[2:5]
	printSlice2("md", md)

	fmt.Println("切片可包含任何类型，甚至包括其它的切片。")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println("向切片追加元素")
	var ps []int
	printSlice(ps)
	if ps == nil {
		fmt.Println("yes")
	}
	ps = append(ps, 1)
	printSlice(ps)
	ps = append(ps, 2)
	printSlice(ps)

	fmt.Println("range")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	npow := make([]int, 10)
	for i := range pow {
		npow[i] = 1 << uint(1)
	}

	for _, v := range npow {
		fmt.Printf("%d ", v)
	}
	//pic.Show(Pic)

}

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			res[i] = append(res[i], uint8(rand.Int()))
		}
	}
	return res
}
