package main

import "fmt"

func main() {
	type point struct {
		x int
		y int
	}

	var p0 point
	var p1 = point{1, 2}

	p0.x = 3
	p0.y = 4

	fmt.Println(p0, p1)

	var m map[int]string

	if m == nil {
		fmt.Println("m is nil")
	}
}
