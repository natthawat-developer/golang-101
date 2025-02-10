package main

import (
    "fmt"
    "time"
)

func sendData(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "Data received!"
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // เริ่ม goroutines
    go sendData(ch1)
    go sendData(ch2)

    // ใช้ select เพื่อรอรับข้อมูลจากทั้งสอง channel
    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    case <-time.After(3 * time.Second): // หากไม่มีการส่งข้อมูลภายใน 3 วินาที
        fmt.Println("Timeout!")
    }
}
