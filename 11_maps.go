package main

import "fmt"

func main() {
    // สร้าง map ที่มี key เป็น string และ value เป็น int
    scores := map[string]int{
        "Alice": 90,
        "Bob":   80,
        "Carol": 85,
    }

    // การเข้าถึงค่าใน map
    fmt.Println("Alice's score:", scores["Alice"])

    // การเพิ่มข้อมูลใหม่
    scores["Dave"] = 95

    // การเช็คว่า key มีอยู่ใน map หรือไม่
    if score, exists := scores["Eve"]; exists {
        fmt.Println("Eve's score:", score)
    } else {
        fmt.Println("Eve's score not found.")
    }

    // การลบข้อมูลจาก map
    delete(scores, "Bob")
    fmt.Println("After deleting Bob:", scores)

    // การวนลูป map
    for name, score := range scores {
        fmt.Println(name, "scored", score)
    }
}
