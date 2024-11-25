package main

import (
	"fmt"
	"sync"
)

// Трогаем goрутины  по уроку
// https://www.youtube.com/watch?v=Qf24zMzMmgI
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fmt.Println(i)
	}

}
