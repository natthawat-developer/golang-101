package main

import "fmt"

// การใช้งานประเภทข้อมูล (Data Types) ใน Go
func main() {
    // 1. ตัวเลขจำนวนเต็ม (Integer)
    var intNum int = 100
    var uintNum uint = 200
    var int8Num int8 = 127  // int8: -128 ถึง 127
    var int16Num int16 = 32767 // int16: -32768 ถึง 32767
    var int32Num int32 = 2147483647 // int32: -2147483648 ถึง 2147483647
    var int64Num int64 = 9223372036854775807 // int64: -9223372036854775808 ถึง 9223372036854775807

    fmt.Println("Integer Values:", intNum, uintNum, int8Num, int16Num, int32Num, int64Num)

    // 2. ตัวเลขทศนิยม (Floating Point)
    var float32Num float32 = 3.14 // มีความแม่นยำ 7 หลัก
    var float64Num float64 = 3.1415926535 // มีความแม่นยำ 15 หลัก

    fmt.Println("Floating Point Values:", float32Num, float64Num)

    // 3. ตัวอักษรและสตริง (Character & String)
    var char rune = 'A' // ใช้ rune เก็บตัวอักษรแบบ Unicode
    var text string = "Hello, Go!"

    fmt.Println("Character and String:", string(char), text)

    // 4. ค่าความจริง (Boolean)
    var isActive bool = true
    var isRunning bool = false

    fmt.Println("Boolean Values:", isActive, isRunning)
}
