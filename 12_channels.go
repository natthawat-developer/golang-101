package main

import (
	"fmt"
	"time"
)

// ฟังก์ชันที่ใช้ channel ในการส่งข้อมูล
func sendData(ch chan string) {
	// ส่งข้อมูลไปยัง channel
	ch <- "Hello from goroutine!"
}

func main() {
	// สร้าง channel สำหรับส่งข้อมูล
	ch := make(chan string)

	// เริ่มต้น goroutine ที่จะส่งข้อมูลไปยัง channel
	go sendData(ch)

	// รับข้อมูลจาก channel
	message := <-ch

	// แสดงข้อมูลที่รับมาจาก channel
	fmt.Println("Received:", message)

	// ใช้เวลาหน่อยเพื่อให้โปรแกรมทำงานเสร็จ
	time.Sleep(1 * time.Second)
}
