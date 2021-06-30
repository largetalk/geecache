package lc

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func Test() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	if res != nil {
		fmt.Println(res)
	}
}