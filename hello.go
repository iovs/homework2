package main

import "fmt"

const (
	a = 1
	b
	c
	d = 2
	e
)

func main() {
	fmt.Println("Hello, 世界, 455566666")
	fmt.Println(a, b, c, d, e) // 1, 1, 1, 2, 2
	fmt.Println("Hello 2")
}
