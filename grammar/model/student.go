package model

import "fmt"

type Student struct {
	Name string
	Age  int
}

//结构体可以使用匿名结构体的任意字段（大写或者小写均可使用）或者方法
type Graduate struct {
	Student
	Gender string
}
type Pupil struct {
	Student
	Score float64
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (p *Pupil) Test() {
	fmt.Println("Pupils are testing")
}

func (g *Graduate) Test() {
	fmt.Println("Graduates are testing")
}
