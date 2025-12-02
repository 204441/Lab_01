package src

import (
	"fmt"

	"lab_01/src"
)

// แก้ YOUR_MODULE_NAME ให้ตรงกับ go.mod ของคุณ
// ตัวอย่างถ้า go.mod เขียนว่า:
//   module github.com/cmu-cs/lab01-hello-world-sync
// ก็ต้อง import:
//   "github.com/cmu-cs/lab01-hello-world-sync/hw01"

func main() {
	max := 10

	lines := src.HelloWorldSync(max)

	for _, line := range lines {
		fmt.Println(line)
	}
}
