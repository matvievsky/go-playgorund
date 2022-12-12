package main

import "fmt"

func main() {
	slice := make([]int, 0, 4)
	slice = append(slice, 1, 2, 3)
	s1 := append(slice, 4)
	s2 := append(slice, 5)
	fmt.Println(s1, s2)
}
