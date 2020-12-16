//结构体的操作练习
// struct和其他语言的class具有同等地位
//golang没有extends继承是通过匿名字段
//面向对象编程=>面向接口编程
package grammar_test

import (
	"github.com/luckylhb90/golang/grammar/model"
	"testing"
)

func Test_struct_1(t *testing.T) {

	user := model.User{
		Name:   "Momo",
		Age:    27,
		Gender: "male",
	}

	println(user.Gender)

}

//test()和UserController绑定
//test()只能由UserController类型的变量调用
