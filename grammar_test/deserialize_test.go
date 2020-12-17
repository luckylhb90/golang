//Tips:反序列化的数据类型应该和原来序列化前的数据类型一致
package grammar_test

import (
	"encoding/json"
	"fmt"
	"github.com/luckylhb90/golang/grammar/model"
	"log"
	"testing"
)

func TestDeserializeStruct(t *testing.T) {
	jsonStr := `{"Name":"dsds","Age":19,"Birthday":"2021-02-03","Skill":"coding"}`
	var jsonDe model.JsonDeserialize
	err := json.Unmarshal([]byte(jsonStr), &jsonDe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonStr)
	fmt.Println(jsonDe)
}

func TestDeserializeSlice(t *testing.T) {
	jsonStr := `[{"Age":17,"Birthday":"2020","Name":"hhe","Skill":"code"},{"Age":32,"Birthday":"2030","Name":"jj","Skill":"eat"}]`
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

//反序列化不需要make，make已经集成到了Unmarshal中
func TestDeserializeMap(t *testing.T) {
	jsonStr := `{"Name":"dsds","Age":19,"Birthday":"2021-02-03","Skill":"coding"}`
	var a map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}
