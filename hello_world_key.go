package src

import (
    "fmt"
    "sync"
)

type Shared struct {
    count int
    turn  string
    out   []string
}

func say(
    word, other string,
    max int,
    s *Shared,
    mu *sync.Mutex,
    wg *sync.WaitGroup,
    done *bool,
) {
    defer wg.Done()

    for {
        mu.Lock()

        if *done {
            mu.Unlock()
            return
        }

        if s.turn != word {
            mu.Unlock()
            continue
        }

        s.count++
        line := fmt.Sprintf("%d %s", s.count, word)
        s.out = append(s.out, line)

        if s.count >= max {
            *done = true
            mu.Unlock()
            return
        }

        s.turn = other
        mu.Unlock()
    }
}

func HelloWorldSync(max int) []string {
    if max <= 0 {
        return []string{}
    }

    s := Shared{
        count: 0,
        turn:  "hello",
        out:   make([]string, 0, max),
    }

    var (
        mu   sync.Mutex
        wg   sync.WaitGroup
        done bool
    )

    wg.Add(2)
    go say("hello", "world", max, &s, &mu, &wg, &done)
    go say("world", "hello", max, &s, &mu, &wg, &done)

    wg.Wait()
    return s.out
}
