package main

import (
	"encoding/json"
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"log"
)

//序列化结构体
func testStruct() {
	m := model.JsonSerilize{
		Name:     "dasd",
		Age:      18,
		Birthday: "2001-02-03",
		Skill:    "coding",
	}
	jsonText, error := json.Marshal(&m)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(jsonText))
}

//序列化map
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "leslie"
	a["age"] = "17"
	a["birthday"] = "2019-02-27"
	a["skill"] = "code"

	jsonString, error := json.Marshal(&a)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(jsonString))
}

//序列化切片
func testSlice() {
	var b []map[string]interface{}
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "leslie"
	a["age"] = "17"
	a["birthday"] = "2019-02-27"
	a["skill"] = "code"
	b = append(b, a)
	var c map[string]interface{}
	c = make(map[string]interface{})
	c["name"] = "Jim"
	c["age"] = "27"
	c["birthday"] = "2010-02-27"
	c["skill"] = "eat"
	b = append(b, c)
	jsonString, error := json.Marshal(&b)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(jsonString))

}

//主函数
func main() {
	//将结构体，map，切片进行序列化
	//testStruct()
	//testMap()
	//testSlice()

}
