//结构体的操作练习
// struct和其他语言的class具有同等地位
//golang没有extends继承是通过匿名字段
//面向对象编程=>面向接口编程
package grammar_test

import (
	"fmt"
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
	//修改结构体字段的value
	user.Age = 30
	fmt.Println(user)

}

func Test_struct_2(t *testing.T) {
	cat1 := model.Cat{
		Name:  "xiaobai",
		Age:   10,
		Color: "baise",
	}
	cat2 := model.Cat{
		Name:  "xiaohei",
		Age:   20,
		Color: "heise",
	}
	fmt.Println(cat1)
	fmt.Println(cat2)

}

//struct中包含array
func Test_struct_3(t *testing.T) {
	var user model.User
	user.Score = [3]int{100, 100, 100}
	user.Name = "hehe"
	fmt.Println(user)

	var p1 model.Struct
	fmt.Println(p1)
	if p1.Per == nil {
		fmt.Println(true)
	}
	//struct中包含slice
	p1.Score = append(p1.Score, 149.9)
	fmt.Println(p1)
	p1.Score = make([]float64, 10)
	p1.Score[0] = 10
	fmt.Println(p1)

	//struct中包含map
	p1.Friend = make(map[string]string)
	p1.Friend["Name"] = "NC"
	fmt.Println(p1)

}

//结构体为值类型，默认为值拷贝
//创造结构体变量的四种方式
func Test_struct_4(t *testing.T) {
	var U1 model.User
	fmt.Println("1---", U1)
	var U2 model.User = model.User{}
	fmt.Println("2---", U2)
	U3 := model.User{}
	fmt.Println("3---", U3)
	fmt.Println("3.1---", &U3)
	fmt.Println("3.2---", *&U3)

	var U4 *model.User = new(model.User)
	fmt.Println("4---", U4)
	fmt.Println("4.1---", &U4)
	fmt.Println("4.2---", *U4)

	var person *model.User = &model.User{
		"name",
		19,
		"male",
		[3]int{1, 2, 3},
	}
	fmt.Println("5---", person)

	(*U4).Name = "hehe"
	(*U4).Score = [...]int{100, 20, 30}
	fmt.Println("6---", *U4)

	// 简化
	U1.Gender = "male"
	fmt.Println("7---", U1)

}

//test()和UserController绑定
//test()只能由UserController类型的变量调用
