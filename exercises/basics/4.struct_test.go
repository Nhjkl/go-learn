package main

import (
	"fmt"
	"testing"
)

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 接收值
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// 值接收者方法（不能修改结构体）
func (p Person) IncrementAgeWrong() {
	p.Age++ // 这不会修改原始结构体
}

// 指针接收者方法（可以修改结构体）
func (p *Person) IncrementAge() {
	p.Age++
}

// 指针接收者方法：修改信息
func (p *Person) ChangeName(name string) {
	p.Name = name
}

type Employee struct {
	Person // 匿名字段
	ID     string
}

// Employee自己的方法
func (e Employee) GetEmployeeInfo() string {
	return fmt.Sprintf("ID: %s, %s", e.ID, e.GetInfo())
}

// 接口示例
type Speaker interface {
	Speak() string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hi, I'm %s", p.Name)
}

// 多态示例
func introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func TestStruct(t *testing.T) {
	t.Log("--------- struct ---------")

	p1 := Person{
		Name: "sean",
		Age:  18,
	}

	t.Logf("p: %v", p1)

	// 简单写法
	p2 := Person{"stone", 18}

	t.Logf("p: %v", p2)

	// 部分初始化
	p3 := Person{Name: "sean"}

	t.Logf("p: %v", p3)

	p1.IncrementAgeWrong()

	t.Log(p1.GetInfo())
	p2.IncrementAge()
	t.Log(p2.GetInfo())
	p3.ChangeName("liuxian")
	t.Log(p3.GetInfo())

	e := Employee{
		Person: Person{"sean", 32},
		ID:     "E001",
	}

	t.Logf("%v", e)
	t.Logf("%v", e.ID)
	t.Logf("%v", e.GetInfo())
	t.Logf("%v", e.GetEmployeeInfo())
	t.Logf("%v", e.Speak())
	t.Logf("%v", p3.Speak())

	introduce(p1)
	introduce(e)
}
