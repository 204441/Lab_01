package src

import (
	"fmt"
	"testing"
)

// ทดสอบว่า HelloWorldSync สร้างลำดับได้ถูกต้อง:
// - ความยาวเท่ากับ max
// - รูปแบบ "N hello/world" ถูกต้อง
func TestHelloWorldSync_Basic(t *testing.T) {
	max := 10
	out := HelloWorldSync(max)

	if len(out) != max {
		t.Fatalf("expected %d lines, got %d", max, len(out))
	}

	for i := 0; i < max; i++ {
		n := i + 1
		expectedWord := "hello"
		if n%2 == 0 {
			expectedWord = "world"
		}
		expected := fmt.Sprintf("%d %s", n, expectedWord)

		if out[i] != expected {
			t.Fatalf("line %d: expected %q, got %q", i, expected, out[i])
		}
	}
}

// ทดสอบกรณี max = 0 (ควรคืน slice ว่าง)
func TestHelloWorldSync_Zero(t *testing.T) {
	out := HelloWorldSync(0)
	if len(out) != 0 {
		t.Fatalf("expected empty slice for max=0, got len=%d", len(out))
	}
}
