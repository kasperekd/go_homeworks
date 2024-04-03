//go:build mytag
// +build mytag

//
//go:generate echo Hello
package main

import (
	"embed"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	var maxNum int
	for i := 0; i < k; i++ {
		maxNum = nums[0]
		maxIndex := 0
		for j := 1; j < len(nums); j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
				maxIndex = j
			}
		}
		nums[maxIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[maxIndex]
		nums = nums[:len(nums)-1]
	}

	return maxNum
}

//go:embed data/*.txt
var files embed.FS

func main() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums1, k1, findKthLargest(nums1, k1))

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums2, k2, findKthLargest(nums2, k2))

	fileInfos, err := files.ReadDir("data")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("List of embedded files:")
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}
