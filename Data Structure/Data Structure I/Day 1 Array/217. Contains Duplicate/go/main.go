package main

import "log"

func containsDuplicate(nums []int) bool {
	type void struct{}
	var member void
	sets := make(map[int]void)
	for _, num := range nums {
		sets[num] = member
	}
	if len(sets) < len(nums) {
		return true
	} else {
		return false
	}
}

func main() {
	nums := []int{1, 2, 3, 1}
	log.Println(containsDuplicate(nums))
	nums = []int{1, 2, 3, 4}
	log.Println(containsDuplicate(nums))
	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	log.Println(containsDuplicate(nums))
}
