package main

import "fmt"

// การประกาศฟังก์ชันแบบพื้นฐาน
func greet(name string) string {
    return "Hello, " + name + "!"
}

// ฟังก์ชันที่มีการคืนค่าหลายค่า
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// ฟังก์ชันที่มีพารามิเตอร์แบบไม่จำกัด (Variadic Function)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// ฟังก์ชันแบบ First-Class (สามารถส่งเป็นพารามิเตอร์ได้)
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    // เรียกใช้งานฟังก์ชัน greet
    fmt.Println(greet("GoLang"))

    // เรียกใช้งานฟังก์ชัน divide
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)

    // เรียกใช้งานฟังก์ชัน sum
    fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

    // ใช้งานฟังก์ชัน applyOperation กับฟังก์ชันนิพจน์ (Anonymous Function)
    result := applyOperation(10, 5, func(x, y int) int { return x + y })
    fmt.Println("Applied Operation Result:", result)
}