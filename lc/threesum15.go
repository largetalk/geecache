package lc

import (
	"fmt"
	"sort"
)

func search(arr []int, n int, idx int) int {
	for i, j := idx, len(arr)-1; i <= j; {
		mid := (i + j) / 2

		if arr[mid] == n {
			return mid
		} else if arr[mid] < n {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return -1
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			k := search(nums, -sum, j+1)
			if k != -1 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

func Test() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))

}
