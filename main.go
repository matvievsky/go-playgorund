package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays(
		[]int{1, 2},
		[]int{3, 4},
	))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	c := make(chan float64, 2)
	f := func(nums []int, c chan float64) {
		length := len(nums)
		i := length / 2
		if length%2 == 0 {
			c <- float64(nums[i-1]+nums[i]) * 0.5
			return
		}
		c <- float64(nums[i])
	}
	go f(nums1, c)
	go f(nums2, c)
	return (<-c + <-c) * 0.5
}
