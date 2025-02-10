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
}