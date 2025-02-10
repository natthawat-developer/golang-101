package main

import "fmt"

// การประกาศตัวแปรแบบต่าง ๆ ใน Go
func main() {
    // 1. การประกาศตัวแปรแบบกำหนดประเภท
    var name string = "John"
    var age int = 25
    var height float64 = 5.9
    var isStudent bool = true

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)

    // 2. การใช้ตัวแปรแบบไม่ต้องระบุประเภท (Type Inference)
    country := "Thailand"
    score := 90.5
    isActive := false

    fmt.Println("Country:", country)
    fmt.Println("Score:", score)
    fmt.Println("Is Active:", isActive)

    // 3. การประกาศตัวแปรหลายตัวพร้อมกัน
    var x, y, z int = 1, 2, 3
    a, b, c := "Go", "Lang", 101

    fmt.Println("x, y, z:", x, y, z)
    fmt.Println("a, b, c:", a, b, c)

    // 4. การใช้ Constants
    const pi float64 = 3.14159
    const appName = "GoLang 101"

    fmt.Println("Pi:", pi)
    fmt.Println("App Name:", appName)
}