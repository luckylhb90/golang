package grammar_test

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/inter"
	"math/rand"
	"sort"
	"testing"
)

func Test_Animal(t *testing.T) {
	d := inter.Dog{}
	x := inter.Tiger{}
	a := inter.Animals{}

	a.DoSomething(d)
	a.DoSomething(x)

}

func Test_Monkey(t *testing.T) {
	var a inter.LittleMonkey
	a.Name = "sun"
	a.Climbing()
	var b inter.Learn = &a
	b.Flying()
	b.Running()
}

func Test_Inter(t *testing.T) {
	var test inter.Test
	var a inter.A = test
	a.SayA()
	a.SayB()
	a.SayC()
}

type TestSlice []inter.Test

func (test TestSlice) Len() int {
	return len(test)
}

func (test TestSlice) Less(i, j int) bool {
	return test[i].Age < test[j].Age
}

func (test TestSlice) Swap2(i, j int) {
	temp := test[i]
	test[i] = test[j]
	test[j] = temp
}
func (test TestSlice) Swap(i, j int) {
	test[i], test[j] = test[j], test[i]
}

func Test_Sort(t *testing.T) {
	slice1 := []int{1, 3, 4, 2, -1, 9}
	sort.Ints(slice1)
	fmt.Println(slice1)

	var slice2 TestSlice
	for i := 0; i < 10; i++ {
		test := inter.Test{
			fmt.Sprintf("ceshi%d", rand.Intn(100)),
			rand.Intn(100),
		}
		slice2 = append(slice2, test)
	}
	fmt.Println(slice2)
	sort.Sort(slice2)
	for _, v := range slice1 {
		fmt.Println(v)
	}

	var x [1]inter.Test

	x[0] = inter.Test{
		"1",
		2,
	}
	fmt.Println(x)
}
