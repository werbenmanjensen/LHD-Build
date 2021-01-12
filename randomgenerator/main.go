package main

import (
	"fmt"
	"time"
)

// rand randomizes from 1 until i
func rand(i int) int{
	if i <= 0 {
		return 0
	}
	return int(time.Now().UnixNano()) % i + 1
}

func main() {
	fmt.Println(rand(10))
}
