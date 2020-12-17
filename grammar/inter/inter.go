package inter

import "fmt"

//golang的interface test
type Animal interface {
	Walk()
	Eat()
}

type Animals struct {
}

type Tiger struct {
}

type Dog struct {
}

func (d Dog) Walk() {
	fmt.Println(" is walking")
}

func (d Dog) Eat() {
	fmt.Println(" is eating")
}

func (t Tiger) Walk() {
	fmt.Println(" is walking")
}

func (t Tiger) Eat() {
	fmt.Println(" is eating")
}

func (a Animals) DoSomething(animal Animal) {
	animal.Eat()
	animal.Walk()
}

//--------
type Monkey struct {
	Name string
}

func (this Monkey) Climbing() {
	fmt.Println(this.Name + "会爬树")
}

type LittleMonkey struct {
	Monkey
}

//不破坏原来老猴子的规则和结构，扩展出新的功能，使用实现接口的方式。
//A结构体继承了B结构体，那么A结构体自动继承了B结构体的字段和方法，并且可以直接使用
//当A结构体需要扩展功能，而又不希望破坏继承关系，则可以采用实现接口的方法

type Learn interface {
	Flying()
	Running()
}

func (m *LittleMonkey) Flying() {
	fmt.Println(m.Name + "is flying")
}

func (m *LittleMonkey) Running() {
	fmt.Println(m.Name + "is running")
}

//----------

// 继承多个接口
// **** 接口是引用类型
type A interface {
	SayA()
	B
	C
}

type B interface {
	SayB()
}

type C interface {
	SayC()
}

type Test struct {
	Name string
	Age  int
}

func (test Test) SayA() {
	fmt.Println("SayA")
}
func (test Test) SayB() {
	fmt.Println("SayB")
}
func (test Test) SayC() {
	fmt.Println("SayC")
}

//----------

type Usb interface {
	Starting()
	Shutdown()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Computer struct {
}

func (this Phone) Starting() {
	fmt.Println("手机启动")
}
func (this Phone) Calling() {
	fmt.Println("手机打电话")
}
func (this Phone) Shutdown() {
	fmt.Println("手机挂断")
}

func (this Computer) Work(usb Usb) {
	usb.Starting()
	if phone, ok := usb.(Phone); ok {
		phone.Calling()
	}
	usb.Shutdown()
}
