package model

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (this Person) Speek() {
	fmt.Println(this.Name + "dafadada")
}

func (this Person) Count() {
	y := 0
	for x := 1; x <= 1000; x++ {
		y += x
	}
	println(y)
}

func (this Person) Count1(n int) {
	y := 0
	for x := 1; x <= n; x++ {
		y += x
	}
	println(y)
}

func (this Person) Sum(x int, y int) {
	fmt.Println(x + y)
}
