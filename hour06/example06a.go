package main

import "fmt"

func main() {
	var nums = make([]string, 3)
	var nums2 = make([]string, 2)
	var nums3 = make([]string, 2)
	nums[0] = "zero"
	nums[1] = "one"
	nums[2] = "two"
	fmt.Println("nums", nums)
	fmt.Println("nums len", len(nums))
	fmt.Println("nums cap", cap(nums))
	copy(nums2, append(nums[:1], nums[1+1:]...))
	fmt.Println("nums2", nums2)
	fmt.Println("nums2 len", len(nums2))
	fmt.Println("nums2 cap", cap(nums2))
	fmt.Println("nums", nums[0:3])
	fmt.Println("nums2", nums2[0:2])
	copy(nums3, nums2[:1])
	fmt.Println("nums2", nums2)
	fmt.Println("nums2 len", len(nums2))
	fmt.Println("nums2 cap", cap(nums2))
	fmt.Println("nums3", nums3)
	fmt.Println("nums3 len", len(nums3))
	fmt.Println("nums3 cap", cap(nums3))
	fmt.Println("nums3", nums2[0:2])
}
