package main

import (
	"fmt"
	"time"
)

// ฟังก์ชันที่ทำงานแบบ concurrent โดยใช้ goroutine
func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond) // หน่วงเวลาเพื่อให้เห็นผลการทำงาน
	}
}

func main() {
	// เริ่มต้น goroutine สำหรับฟังก์ชัน printMessage
	go printMessage("Hello from goroutine!")

	// เรียกใช้งาน printMessage ใน main goroutine
	printMessage("Hello from main!")

	// ใช้เวลาสักครู่เพื่อให้ goroutines ทำงานเสร็จ
	time.Sleep(1 * time.Second)
}
