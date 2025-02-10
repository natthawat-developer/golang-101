package main

import (
	"errors"
	"fmt"
)

// ฟังก์ชันที่ทำงานและอาจเกิดข้อผิดพลาด
func divide(a, b int) (int, error) {
	if b == 0 {
		// คืนค่าข้อผิดพลาดหาก b เป็น 0
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// เรียกฟังก์ชัน divide และตรวจสอบข้อผิดพลาด
	result, err := divide(10, 0)
	if err != nil {
		// หากเกิดข้อผิดพลาด ให้แสดงข้อความผิดพลาด
		fmt.Println("Error:", err)
	} else {
		// หากไม่มีข้อผิดพลาด แสดงผลลัพธ์
		fmt.Println("Result:", result)
	}

	// ตัวอย่างอีกกรณีที่ไม่มีข้อผิดพลาด
	result, err = divide(10, 2)
	if err != nil {
		// หากเกิดข้อผิดพลาด ให้แสดงข้อความผิดพลาด
		fmt.Println("Error:", err)
	} else {
		// หากไม่มีข้อผิดพลาด แสดงผลลัพธ์
		fmt.Println("Result:", result)
	}
}
