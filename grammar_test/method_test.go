package grammar_test

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"testing"
)

func Test_Person(t *testing.T) {
	var person model.Person
	person.Name = "aaa"
	person.Speek()
	person.Count()

	var n int
	fmt.Scanln(&n)

	person.Count1(n)

	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	person.Sum(a, b)
}

func Test_Circle(t *testing.T) {
	var circle = model.Circle{3.0}
	res := circle.Area()
	fmt.Println(res)
}

func Test_Method(t *testing.T) {
	var user model.MethonUtils
	user.Print()
	user.Print2(10, 2)

	area := user.Area(10.4, 2.3)

	fmt.Println(area)
	user.Judge(9)

	var test model.Method
	test.Test(10, 9, "@")
	var x model.Method1
	res := x.Test1(1.0, 2.4, "+")
	fmt.Println(res)
}
