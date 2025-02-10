package main

import "fmt"

// การทำงานกับ Array และ Slice ใน Go
func main() {
    // 1. Array
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", arr)

    // 2. Slice
    slice := []int{4, 5, 6}
    fmt.Println("Slice:", slice)

    // 3. การเพิ่มค่าใน Slice
    slice = append(slice, 7)
    fmt.Println("Slice after append:", slice)

    // 4. การเข้าถึงข้อมูลใน Array และ Slice
    fmt.Println("First element of array:", arr[0])
    fmt.Println("First element of slice:", slice[0])

    // 5. Slice สามารถมีขนาดยืดหยุ่นได้
    subSlice := slice[1:4]
    fmt.Println("Sub-slice:", subSlice)
}
