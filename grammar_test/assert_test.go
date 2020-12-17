package grammar_test

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/inter"
	"testing"
)

//断言类型失败，报错panic，为避免panic错误，我们对程序进行改写
func Test_Assert(t *testing.T) {
	var a interface{}
	var j int64
	a = j
	b := a.(int32)
	fmt.Println(b)
	//断言类型失败，报错panic，为避免panic错误，我们对程序进行改写
}

func Test_Assert2(t *testing.T) {
	var a interface{}
	var j int64
	a = j
	b, ok := a.(int32)
	if !ok {
		fmt.Println("类型断言错误")
		fmt.Println(b)
	}
	c, ok := a.(int64)
	if ok {
		fmt.Println("类型断言成功")
		fmt.Println(c)
	}
}

func Test_Assert3(t *testing.T) {
	var a interface{}
	b := inter.Test{"hhh", 22}
	a = b
	c, ok := a.(inter.Test)
	if ok {
		fmt.Println("断言成功")
		fmt.Println(c)
	} else {
		fmt.Println("断言失败   ")
	}

}
