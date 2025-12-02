package src

import (
	"fmt"
	"sync"
)

// Shared: state ร่วมระหว่าง goroutine "hello" และ "world"
type Shared struct {
	count int      // นับลำดับ 1..max
	turn  string   // "hello" หรือ "world" (ตอนนี้ถึงตาใคร)
	out   []string // เก็บผลลัพธ์เป็น slice ของ string เช่น "1 hello"
}

// say ควรจะ:
//   - ใช้ mutex ป้องกันการแก้ไข s.count, s.turn, s.out, *done
//   - ผลัดกันทำงานตามค่า s.turn
//   - เพิ่ม count ทีละ 1 และบันทึก string "N word" ลงใน s.out
//   - หยุดเมื่อ count >= max แล้วตั้ง *done = true
//
// นักศึกษาต้องเติมส่วน:
//   - ใช้ wg.Done() (แนะนำใช้ defer wg.Done())
//   - ใส่ mu.Lock() / mu.Unlock() ในจุดที่เหมาะสม (critical section)
func say(
	word, other string,
	max int,
	s *Shared,
	mu *sync.Mutex,
	wg *sync.WaitGroup,
	done *bool,
) {
	// ป้องกัน error “unused parameter” ให้ compile ผ่าน
	_ = mu
	_ = wg
	_ = done

	// TODO: นักศึกษา: ใส่ defer wg.Done() ที่นี่
	// defer wg.Done()

	for {
		// TODO: นักศึกษา: ก่อนเข้าถึง shared state ให้ lock
		// mu.Lock()

		// ถ้าทำงานเสร็จแล้ว (อีกฝั่งตั้ง done ไว้) ให้ออก
		if *done {
			// TODO: นักศึกษา: ปลด lock ก่อน return
			// mu.Unlock()
			return
		}

		// ถ้ายังไม่ถึงตาของ goroutine นี้ ให้ปลด lock แล้ววนใหม่
		if s.turn != word {
			// TODO: นักศึกษา: ปลด lock ก่อน continue
			// mu.Unlock()
			continue
		}

		// ---- critical section ----
		s.count++
		line := fmt.Sprintf("%d %s", s.count, word)
		s.out = append(s.out, line)

		// ถ้าครบ max แล้ว ให้ตั้ง done = true และออกจาก goroutine นี้
		if s.count >= max {
			*done = true
			// TODO: นักศึกษา: ปลด lock ก่อน return
			// mu.Unlock()
			return
		}

		// ส่งตาต่อไปให้อีกคำหนึ่ง
		s.turn = other

		// TODO: นักศึกษา: ปลด lock ก่อนจะวนลูปใหม่
		// mu.Unlock()
	}
}

// HelloWorldSync คือฟังก์ชันหลักที่ test จะเรียกใช้
//
// ต้อง:
//   - ใช้ goroutine 2 ตัว (hello/world)
//   - ใช้ mutex + waitgroup เพื่อให้ผลลัพธ์ถูกต้องและไม่มี data race
//   - คืนค่า []string ในรูป:
//       []string{"1 hello", "2 world", "3 hello", ...}
func HelloWorldSync(max int) []string {
	if max <= 0 {
		return []string{}
	}

	// เตรียม shared state เริ่มต้น
	s := Shared{
		count: 0,
		turn:  "hello",          // ให้ "hello" เริ่มก่อน
		out:   make([]string, 0, max),
	}

	var (
		mu   sync.Mutex
		wg   sync.WaitGroup
		done bool
	)

	// ป้องกัน unused variable ให้ compile ผ่าน
	_ = mu
	_ = wg
	_ = done

	// TODO: นักศึกษา: บอก WaitGroup ว่าจะมี goroutine กี่ตัว (2 ตัว)
	// wg.Add(2)

	// สร้าง goroutine สองตัว
	go say("hello", "world", max, &s, &mu, &wg, &done)
	go say("world", "hello", max, &s, &mu, &wg, &done)

	// TODO: นักศึกษา: รอให้ goroutine ทั้งสองตัวทำงานเสร็จ
	// wg.Wait()

	// เมื่อทุกอย่างเสร็จแล้ว คืน slice ของผลลัพธ์
	return s.out
}
