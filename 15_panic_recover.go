package main

import "fmt"

// ฟังก์ชันที่ใช้ panic เพื่อสร้างข้อผิดพลาด
func causePanic() {
	fmt.Println("Going to panic!")
	panic("Something went wrong!")
}

// ฟังก์ชันที่ใช้ recover เพื่อจับข้อผิดพลาดจาก panic
func handlePanic() {
	// ใช้ defer เพื่อเรียก recover และจับ panic
	defer func() {
		if r := recover(); r != nil {
			// recover จะคืนค่าข้อผิดพลาดที่ถูก panic
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// เรียกฟังก์ชัน causePanic ที่จะทำให้เกิด panic
	causePanic()
}

func main() {
	// เรียก handlePanic ที่จะทำการ recover จาก panic
	handlePanic()

	// โค้ดหลังจากการ recover ยังคงทำงานได้ตามปกติ
	fmt.Println("Program continues after recovery.")
}
