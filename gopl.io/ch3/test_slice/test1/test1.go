package main

import "fmt"

func main() {
	s1 := [...]int{1, 2, 3, 4, 5}
	x := s1[:len(s1)-1]
	y := s1[:0]

	fmt.Println(x)
	fmt.Println(y)
}
