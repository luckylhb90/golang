package grammar_test

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"testing"
)

//使用工厂模式操作model
func Test_factory(t *testing.T) {
	fmt.Println("golang使用工厂模式！")
	student := model.New("Leslie", 120)
	fmt.Println(student)

}
