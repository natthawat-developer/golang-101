package main

import (
	"testing"
)

// ฟังก์ชันที่ต้องการทดสอบ
func Add(a, b int) int {
	return a + b
}

// ฟังก์ชันทดสอบสำหรับฟังก์ชัน Add
func TestAdd(t *testing.T) {
	// ทดสอบการบวกเลข
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}

	// ทดสอบกรณีบวกเลขลบ
	result = Add(-1, -1)
	if result != -2 {
		t.Errorf("Add(-1, -1) = %d; want -2", result)
	}

	// ทดสอบกรณีบวกศูนย์
	result = Add(0, 0)
	if result != 0 {
		t.Errorf("Add(0, 0) = %d; want 0", result)
	}
}
