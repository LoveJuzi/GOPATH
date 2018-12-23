package main

import (
	"fmt"
	"time"
)

func main() {
	var name = "Eric"
	go func(name string) {
		fmt.Println("Hello,", name)
	}(name)
	name = "Harry"
	time.Sleep(time.Millisecond)
}
