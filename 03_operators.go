package main

import "fmt"

// การใช้งานตัวดำเนินการ (Operators) ใน Go
func main() {
    // 1. ตัวดำเนินการทางคณิตศาสตร์ (Arithmetic Operators)
    a, b := 10, 3
    fmt.Println("Arithmetic Operators:")
    fmt.Println("Addition:", a, "+", b, "=", a+b)
    fmt.Println("Subtraction:", a, "-", b, "=", a-b)
    fmt.Println("Multiplication:", a, "*", b, "=", a*b)
    fmt.Println("Division:", a, "/", b, "=", a/b)
    fmt.Println("Modulus:", a, "%", b, "=", a%b)

    // 2. ตัวดำเนินการเปรียบเทียบ (Comparison Operators)
    fmt.Println("\nComparison Operators:")
    fmt.Println(a, "==", b, "->", a == b)
    fmt.Println(a, "!=" , b, "->", a != b)
    fmt.Println(a, ">", b, "->", a > b)
    fmt.Println(a, "<", b, "->", a < b)
    fmt.Println(a, ">=", b, "->", a >= b)
    fmt.Println(a, "<=", b, "->", a <= b)

    // 3. ตัวดำเนินการทางตรรกะ (Logical Operators)
    x, y := true, false
    fmt.Println("\nLogical Operators:")
    fmt.Println("AND:", x, "&&", y, "->", x && y)
    fmt.Println("OR:", x, "||", y, "->", x || y)
    fmt.Println("NOT:", "!x ->", !x)

    // 4. ตัวดำเนินการบิต (Bitwise Operators)
    p, q := 5, 3 // 5 = 101, 3 = 011 (ในเลขฐาน 2)
    fmt.Println("\nBitwise Operators:")
    fmt.Println("AND:", p, "&", q, "->", p&q)
    fmt.Println("OR:", p, "|", q, "->", p|q)
    fmt.Println("XOR:", p, "^", q, "->", p^q)
    fmt.Println("Left Shift:", p, "<< 1 ->", p<<1)
    fmt.Println("Right Shift:", p, ">> 1 ->", p>>1)

    // 5. ตัวดำเนินการกำหนดค่า (Assignment Operators)
    fmt.Println("\nAssignment Operators:")
    num := 10
    fmt.Println("num =", num)
    num += 5
    fmt.Println("num += 5 ->", num)
    num -= 3
    fmt.Println("num -= 3 ->", num)
    num *= 2
    fmt.Println("num *= 2 ->", num)
    num /= 4
    fmt.Println("num /= 4 ->", num)
    num %= 3
    fmt.Println("num %= 3 ->", num)
}
