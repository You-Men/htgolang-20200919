package main

import "fmt"

/*
针对问题2的切片, 将最大的数字移动到切片的最后一位
原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
*/

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}
	var maxIndex int
	for x, v := range nums {
		if v > nums[maxIndex] {
			maxIndex = x
		}
	}

	var sortNums = []int{}
	for i, v := range nums {
		if i != maxIndex {
			sortNums = append(sortNums, v)
		}
	}

	sortNums = append(sortNums, nums[maxIndex])
	fmt.Println(sortNums)
}
