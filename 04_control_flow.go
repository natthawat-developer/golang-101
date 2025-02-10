package main

import "fmt"

// การควบคุมโฟลว์ของโปรแกรม (Control Flow) ใน Go
func main() {
    // 1. คำสั่ง if-else
    num := 10
    if num > 0 {
        fmt.Println("num เป็นค่าบวก")
    } else if num < 0 {
        fmt.Println("num เป็นค่าลบ")
    } else {
        fmt.Println("num เป็นศูนย์")
    }

    // 2. คำสั่ง switch
    day := 3
    switch day {
    case 1:
        fmt.Println("วันจันทร์")
    case 2:
        fmt.Println("วันอังคาร")
    case 3:
        fmt.Println("วันพุธ")
    default:
        fmt.Println("วันอื่น ๆ")
    }

    // 3. วนลูป for
    fmt.Println("\nวนลูป for แบบปกติ")
    for i := 1; i <= 5; i++ {
        fmt.Println("รอบที่", i)
    }

    // 4. วนลูป while-style ด้วย for
    fmt.Println("\nวนลูป while-style")
    count := 0
    for count < 3 {
        fmt.Println("Count:", count)
        count++
    }

    // 5. วนลูป range ใช้กับ array/slice/map
    fmt.Println("\nวนลูป range")
    nums := []int{10, 20, 30, 40}
    for index, value := range nums {
        fmt.Println("Index:", index, "Value:", value)
    }

    // 6. การใช้ goto
    fmt.Println("\nการใช้ goto")
    num2 := 5
    if num2 > 0 {
        fmt.Println("num2 เป็นค่าบวก")
        goto skipMessage
    }
    fmt.Println("ข้อความนี้จะไม่แสดง")
skipMessage:
    fmt.Println("โปรแกรมข้ามไปที่นี่")

    // 7. การใช้ defer, panic, และ recover
    fmt.Println("\nการใช้ defer, panic, และ recover")
    defer fmt.Println("สิ้นสุดการทำงานฟังก์ชัน") // defer ใช้สำหรับเลื่อนการทำงานไปจนกว่าฟังก์ชันจะสิ้นสุด

    // ใช้ panic เพื่อทำให้โปรแกรมหยุดทำงาน
    fmt.Println("โปรแกรมเริ่มทำงาน")
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("เกิดข้อผิดพลาด:", r)
        }
    }() // ใช้ recover เพื่อจับข้อผิดพลาดจาก panic
    panic("เกิดข้อผิดพลาดในโปรแกรม")
    fmt.Println("ข้อความนี้จะไม่ถูกแสดง")
}
