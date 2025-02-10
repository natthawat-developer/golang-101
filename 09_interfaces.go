package main

import "fmt"

// การประกาศ interface
type Greeter interface {
	Greet() string
}

// การประกาศ struct ที่จะ implement interface
type Person struct {
	Name string
}

// ฟังก์ชันที่ implement interface `Greeter`
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Name string
}

// ฟังก์ชันที่ implement interface `Greeter` สำหรับ Dog
func (d Dog) Greet() string {
	return "Woof! I am a dog named " + d.Name
}

func main() {
	// สร้าง instance ของ struct ที่ implement interface
	person := Person{Name: "Natthawat"}
	dog := Dog{Name: "Buddy"}

	// ฟังก์ชันรับประเภทข้อมูลที่เป็น interface
	printGreeting(person)
	printGreeting(dog)
}

// ฟังก์ชันที่รับ interface
func printGreeting(g Greeter) {
	fmt.Println(g.Greet())
}
