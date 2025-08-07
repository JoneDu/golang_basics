package oop

import (
	"fmt"
	"math"
)

//✅面向对象
//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格。

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.length + 2*r.width
}

func GetAreaAndPerimeter() {
	rectangle := Rectangle{length: 4.5, width: 2.0}
	circle := Circle{radius: 4.5}

	fmt.Printf("rectangle.Perimeter(): %+v\n", rectangle.Perimeter())
	fmt.Printf("rectangle.Area(): %+v\n", rectangle.Area())
	fmt.Printf("circle.Perimeter(): %+v\n", circle.Perimeter())
	fmt.Printf("circle.Area(): %+v\n", circle.Area())
}

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) printInfo() {
	fmt.Printf("e: %+v\n", *e)
}

func PrintEmployeeInfo() {
	// 初始化Employee
	emp := Employee{
		Person: Person{
			Name: "Bruce",
			Age:  29,
		},
		EmployeeID: 1002,
	}
	emp.printInfo()
}
