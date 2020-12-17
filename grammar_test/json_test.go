package grammar_test

import (
	"encoding/json"
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"log"
	"testing"
)

//序列化结构体
func TestStruct(t *testing.T) {
	m := model.JsonSerilize{
		"dsds",
		19,
		"2021-02-03",
		"coding",
	}
	fmt.Println(m)
	fmt.Println(&m)
	jsonText, error := json.Marshal(m)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(jsonText)
	fmt.Println(string(jsonText))
}

//序列化map
func TestMap(t *testing.T) {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "hehe"
	a["age"] = "17"
	a["birthday"] = "2020"
	a["skill"] = "coding"

	jsonTxt, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonTxt))
}

//序列化切片
func Test_Slice(t *testing.T) {
	var b []map[string]interface{}
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "hhe"
	a["age"] = 17
	a["birthday"] = "2020"
	a["skill"] = "code"
	b = append(b, a)

	c := make(map[string]interface{})
	c["name"] = "jj"
	c["age"] = 32
	c["birthday"] = "2030"
	c["skill"] = "eat"
	b = append(b, c)

	js, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(js))
}
