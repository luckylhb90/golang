package grammar_test

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"testing"
)

//继承的用法
//代码复用性提高
//可维护性和可扩展性提高了

func Test_extends(t *testing.T) {

	p := model.Pupil{model.Student{"hehe", 30}, 23.2}
	g := model.Graduate{model.Student{"ddd", 2}, "male"}

	p.Test()
	p.Student.SetAge(20)
	fmt.Println(p)

	age1 := p.Student.GetAge()
	fmt.Println(age1)

	g.Test()
	g.Student.SetAge(21)

	age2 := g.Student.GetAge()
	fmt.Println(age2)

}
