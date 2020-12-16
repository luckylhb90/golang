//map引用类型
package grammar

import (
	"fmt"
	"testing"
)

func Test_map_1(t *testing.T) {
	var test map[string]string
	test = make(map[string]string)
	test["name"] = "Momo"
	test["age"] = "8"
	out(test)

	test["skill"] = "coding"
	out(test)

	delete(test, "age")

	out(test)

	for i, _ := range test {
		delete(test, i)
	}
	out(test)

}

func Test_map_2(t *testing.T) {

	test := make(map[string]string)
	test["name"] = "Momo"
	test["age"] = "8"
	out(test)
	test = make(map[string]string)
	out(test)

}

func Test_map_3(t *testing.T) {
	var slice []map[string]string
	slice = make([]map[string]string, 1)
	test := make(map[string]string)
	test["name"] = "Momo"
	test["age"] = "8"
	out(test)
	slice = append(slice, test)
	fmt.Println(slice)
}

func out(tests map[string]string) (int, error) {
	return fmt.Println(tests)
}
