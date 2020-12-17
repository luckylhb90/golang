package model

type JsonSerilize struct {
	//使用tag设置别名，反射机制
	Name     string `json:"mingzi"`
	Age      int
	Birthday string
	Skill    string
}

//反序列化
//将json反序列化成切片，map,struct
type JsonDeserialize struct {
	Name     string
	Age      int
	Birthday string
	Skill    string
}

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
