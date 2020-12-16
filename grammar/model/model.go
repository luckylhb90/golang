package model

type User struct {
	Name   string
	Age    int
	Gender string
	Score  [3]int
}

type Cat struct {
	Name     string
	Color    string
	Age      int
	Behavior string
}
type Struct struct {
	Per    *int
	Name   string
	Age    int
	Score  []float64
	Hobby  [3]float64
	Friend map[string]string
}

type Student struct {
	Name string
	Age  int
}

//适用于结构体首字母小写的情况，可使用工厂模式
func New(name string, age int) *Student {
	//return &Student{
	//	Name: name,
	//	Age:  age,
	//}
	return &Student{
		Name: name,
		Age:  age,
	}
}
