package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestHelloWorldSync_Basic ทดสอบว่าผลลัพธ์มีจำนวนและลำดับที่ถูกต้อง
func TestHelloWorldSync_Basic(t *testing.T) {
	max := 10
	result := HelloWorldSync(max)

	// ทดสอบจำนวน
	if len(result) != max {
		t.Errorf("ผลลัพธ์ควรมี %d บรรทัด แต่ได้ %d บรรทัด", max, len(result))
		t.Logf("ผลลัพธ์ที่ได้: %v", result)
		return
	}

	// ทดสอบลำดับและรูปแบบ
	expected := []string{
		"1 hello",
		"2 world",
		"3 hello",
		"4 world",
		"5 hello",
		"6 world",
		"7 hello",
		"8 world",
		"9 hello",
		"10 world",
	}

	for i, expectedLine := range expected {
		if result[i] != expectedLine {
			t.Errorf("บรรทัดที่ %d: คาดหวัง %q แต่ได้ %q", i+1, expectedLine, result[i])
		}
	}
}

// TestHelloWorldSync_Alternating ทดสอบว่า hello และ world สลับกันอย่างถูกต้อง
func TestHelloWorldSync_Alternating(t *testing.T) {
	max := 20
	result := HelloWorldSync(max)

	if len(result) != max {
		t.Fatalf("ผลลัพธ์ควรมี %d บรรทัด แต่ได้ %d บรรทัด", max, len(result))
	}

	// ตรวจสอบว่าสลับกันระหว่าง hello และ world
	for i, line := range result {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			t.Errorf("บรรทัดที่ %d มีรูปแบบไม่ถูกต้อง: %q (ควรเป็น \"N word\")", i+1, line)
			continue
		}

		word := parts[1]

		// บรรทัดที่ 1, 3, 5, ... (index 0, 2, 4, ...) ควรเป็น "hello"
		// บรรทัดที่ 2, 4, 6, ... (index 1, 3, 5, ...) ควรเป็น "world"
		if i%2 == 0 && word != "hello" {
			t.Errorf("บรรทัดที่ %d: คาดหวัง \"hello\" แต่ได้ %q", i+1, word)
		}
		if i%2 == 1 && word != "world" {
			t.Errorf("บรรทัดที่ %d: คาดหวัง \"world\" แต่ได้ %q", i+1, word)
		}

		// ตรวจสอบหมายเลข
		expectedNum := fmt.Sprintf("%d", i+1)
		if parts[0] != expectedNum {
			t.Errorf("บรรทัดที่ %d: คาดหวังหมายเลข %s แต่ได้ %q", i+1, expectedNum, parts[0])
		}
	}
}

// TestHelloWorldSync_Format ทดสอบรูปแบบของแต่ละบรรทัด
func TestHelloWorldSync_Format(t *testing.T) {
	max := 5
	result := HelloWorldSync(max)

	if len(result) != max {
		t.Fatalf("ผลลัพธ์ควรมี %d บรรทัด แต่ได้ %d บรรทัด", max, len(result))
	}

	for i, line := range result {
		// ตรวจสอบว่ามี format "N word"
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			t.Errorf("บรรทัดที่ %d: รูปแบบไม่ถูกต้อง %q (ควรเป็น \"หมายเลข คำ\")", i+1, line)
			continue
		}

		// ตรวจสอบว่าคำเป็น hello หรือ world เท่านั้น
		word := parts[1]
		if word != "hello" && word != "world" {
			t.Errorf("บรรทัดที่ %d: คำไม่ถูกต้อง %q (ควรเป็น \"hello\" หรือ \"world\")", i+1, word)
		}
	}
}

// TestHelloWorldSync_Zero ทดสอบกรณีขอบเขต: max = 0
func TestHelloWorldSync_Zero(t *testing.T) {
	result := HelloWorldSync(0)
	if len(result) != 0 {
		t.Errorf("เมื่อ max=0 ควรได้ slice ว่าง แต่ได้ %v", result)
	}
}

// TestHelloWorldSync_Negative ทดสอบกรณีขอบเขต: max < 0
func TestHelloWorldSync_Negative(t *testing.T) {
	result := HelloWorldSync(-5)
	if len(result) != 0 {
		t.Errorf("เมื่อ max<0 ควรได้ slice ว่าง แต่ได้ %v", result)
	}
}
