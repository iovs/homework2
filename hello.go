package main

import "fmt"

const (
	a = 1
	b
	c
	d = 2
	e
	f
	g = 8
	h
)

func main() {
	fmt.Println("Hello, 世界, 455566666")
	fmt.Println(a, b, c, d, e, f, g, h) // 1, 1, 1, 2, 2, 2, 8, 8
}
