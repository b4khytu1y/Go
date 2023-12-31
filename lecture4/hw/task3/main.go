package main

import (
	"fmt"
	"sync"
)

func main() {
	m := &sync.Map{}

	var muKey1 sync.Mutex
	var muKey2 sync.Mutex

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			muKey1.Lock()
			value, ok := m.Load("key")
			if !ok {
				m.Store("key", 1)
			} else {
				m.Store("key", value.(int)+1)
			}
			muKey1.Unlock()
		}()
	}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			muKey2.Lock()
			value, ok := m.Load("key2")
			if !ok {
				m.Store("key2", 1)
			} else {
				m.Store("key2", value.(int)+1)
			}
			muKey2.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(m.Load("key"))
	fmt.Println(m.Load("key2"))
}
