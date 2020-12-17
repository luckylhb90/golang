package model

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (this Circle) Area() float64 {
	return this.Radius * this.Radius * math.Pi
}

type MethonUtils struct {
}

type Method struct {
	x     int
	y     int
	shape string
}

type Method1 struct {
	x     float64
	y     float64
	shape string
}

func (user MethonUtils) Print() {
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 8; y++ {
			fmt.Println("*")
		}
		fmt.Println("")
	}
}

func (user MethonUtils) Print2(m int, n int) {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Println("*")
		}
		fmt.Println("")
	}
}

func (user MethonUtils) Area(width float64, length float64) float64 {
	return width * length
}

func (this MethonUtils) Judge(num int) {
	if num%2 == 0 {
		fmt.Printf("%v是偶数", num)
	} else {
		fmt.Printf("%v是奇数", num)
	}
}

func (this Method) Test(m int, n int, shape string) {
	for x := 1; x <= m; x++ {
		for y := 1; y <= n; y++ {
			fmt.Print(shape)
		}
		fmt.Println("")
	}
}

func (this Method1) Test1(m float64, n float64, op string) float64 {
	res := 0.0
	switch op {
	case "+":
		res = m + n
		return res
	case "-":
		res = m - n
		return res
	case "*":
		res = m * n
		return res
	case "/":
		res = m / n
		return res
	default:
		fmt.Println("操作符有误")
		return res
	}
}
