package main

import "fmt"

func main() {
	var nums = make([]string, 3)
	nums[0] = "zero"
	nums[1] = "one"
	nums[2] = "two"
	fmt.Println("nums", nums)
	fmt.Println("nums len", len(nums))
	fmt.Println("nums cap", cap(nums))
	var nums2 = append(nums[:1], nums[1+1:]...)
	fmt.Println("nums2", nums2)
	fmt.Println("nums2 len", len(nums2))
	fmt.Println("nums2 cap", cap(nums2))
	fmt.Println("nums", nums[0:3])
	fmt.Println("nums2", nums2[0:3])
	var nums3 = append(nums, "three")
	fmt.Println("nums3", nums3)
	fmt.Println("nums3 len", len(nums3))
	fmt.Println("nums3 cap", cap(nums3))
	fmt.Println("nums2", nums2[0:3])
	fmt.Println("nums3", nums3[0:6])
	fmt.Println("nums", nums)
	fmt.Println("nums len", len(nums))
	fmt.Println("nums cap", cap(nums))
}
