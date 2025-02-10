package main

import (
	"testing"
)

// ฟังก์ชันที่เราจะทำการ benchmark
func Add(a, b int) int {
	return a + b
}

// ฟังก์ชัน benchmark สำหรับฟังก์ชัน Add
func BenchmarkAdd(b *testing.B) {
	// ทำการรัน benchmark ด้วยการเรียกฟังก์ชัน Add ในรอบหลายๆ ครั้ง
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
