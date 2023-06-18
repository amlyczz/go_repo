package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println("无限循环")
	/*
		for {
		}
	*/

	x := 1
	if x < 0 {
		fmt.Println("x<0")
	} else {
		fmt.Println("x>=0")
	}

	fmt.Println("short if")
	print(pow(3, 2, 10))

	fmt.Println("switch")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoo")
	default:
		fmt.Println("Good evening")
	}
	fmt.Println("defer 语句会将函数推迟到外层函数返回之后执行。")
	fmt.Println("推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。")
	fmt.Println("推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。")
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
	}
	return lim
}
