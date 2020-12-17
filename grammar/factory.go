package main

import (
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
)

//使用工厂模式操作model
func main() {
	fmt.Println("golang使用工厂模式！")
	student := model.New("Leslie", 120)
	fmt.Println(student)

}
