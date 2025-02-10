package main

import "fmt"

// การประกาศ struct
type Person struct {
	Name    string
	Age     int
	Address string
}

// ฟังก์ชันเพื่อแสดงข้อมูลของ Person
func (p Person) Introduce() {
	fmt.Printf("Hi, my name is %s. I am %d years old, and I live at %s.\n", p.Name, p.Age, p.Address)
}

func main() {
	// การสร้าง instance ของ struct
	person1 := Person{
		Name:    "Natthawat",
		Age:     28,
		Address: "Bangkok, Thailand",
	}

	// เรียกใช้ฟังก์ชันที่เป็น method ของ struct
	person1.Introduce()

	// การเข้าถึงและแก้ไขสมาชิกของ struct
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Address:", person1.Address)

	// การสร้าง struct แบบไม่ต้องใช้ชื่อฟิลด์
	person2 := Person{"Sarah", 25, "New York, USA"}
	person2.Introduce()
}
