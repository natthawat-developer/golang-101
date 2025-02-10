package main

import "fmt"

// Closure: ฟังก์ชันที่สามารถเข้าถึงตัวแปรจากภายนอกแม้ว่าไฟล์นั้นจะจบการทำงานไปแล้ว
func main() {
	// ฟังก์ชันภายในที่สามารถเข้าถึงตัวแปรนอก
	increment := func() func() int {
		// ตัวแปรภายนอกที่ถูกเก็บไว้ใน closure
		count := 0
		return func() int {
			count++
			return count
		}
	}

	// สร้าง closure ที่จะทำการเพิ่มค่า
	inc := increment()

	// เรียกใช้งาน closure หลายครั้ง
	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3
}
