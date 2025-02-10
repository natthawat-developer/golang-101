package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex // ประกาศ Mutex เพื่อใช้ในการป้องกัน race condition
)

// ฟังก์ชันที่จะเพิ่มค่า counter โดยใช้ Mutex
func increment() {
	mu.Lock() // ล็อก Mutex ก่อนทำการแก้ไขค่า
	counter++
	mu.Unlock() // ปลดล็อก Mutex หลังจากที่ทำการแก้ไขเสร็จ
}

func main() {
	// ใช้ wait group เพื่อรอการทำงานของ goroutines ทั้งหมดให้เสร็จ
	var wg sync.WaitGroup

	// สร้าง 1000 goroutines เพื่อเพิ่มค่า counter
	for i := 0; i < 1000; i++ {
		wg.Add(1) // เพิ่ม count ของ wait group
		go func() {
			defer wg.Done() // ลด count ของ wait group เมื่อ goroutine ทำงานเสร็จ
			increment()
		}()
	}

	// รอจนกว่าทุก goroutine จะเสร็จการทำงาน
	wg.Wait()

	// แสดงผลลัพธ์ของ counter
	fmt.Println("Final Counter:", counter)
}
