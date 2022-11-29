package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	test(arr)
	fmt.Println(arr)
	slc := make([]int, 0, 3)
	slc = append(slc, 1, 2, 3)
	test(slc)
	fmt.Println(slc)
}

func test(nums []int) {
	nums[1] = 3
}
