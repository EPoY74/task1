package main

import (
	"fmt"
)

// Трогаем goрутины  по уроку
// https://www.youtube.com/watch?v=Qf24zMzMmgI
func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println(i)
	}

}
