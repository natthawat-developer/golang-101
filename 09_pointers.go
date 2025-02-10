package main

import "fmt"

// ฟังก์ชันที่ใช้ pointer เพื่อแก้ไขค่าของตัวแปร
func updateValue(num *int) {
	// การเข้าถึงค่าผ่าน pointer และแก้ไขค่า
	*num = 10
}

func main() {
	// ประกาศตัวแปรและ pointer
	var x int = 5
	var ptr *int = &x // ตัวแปร ptr เก็บที่อยู่ของ x

	// แสดงค่าก่อนการใช้ pointer
	fmt.Println("Before:", x)

	// การใช้ pointer ในการแก้ไขค่าตัวแปร
	updateValue(ptr)

	// แสดงค่าหลังการใช้ pointer
	fmt.Println("After:", x)

	// การแสดงที่อยู่ของตัวแปร
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer ptr points to:", ptr)
	fmt.Println("Value at the address ptr points to:", *ptr)
}
