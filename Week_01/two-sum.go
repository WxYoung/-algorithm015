package main

func main() {
	nums := []int{-10, -1, -18, -19}
	target := -19

	twoSum(nums, target)
	twoSum2(nums, target)
}

func twoSum2(nums []int, target int) []int {
	//hash法
	result := []int{}
	m := make(map[int]int)

	for i, k := range nums {
		if v, exist := m[target-k]; exist {
			result = []int{v, i}
		}
		m[k] = i
	}

	return result
}

func twoSum(nums []int, target int) []int {
	args := make([]int, 0)
	//暴力循环法
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				args = []int{i, j}
			}
		}
	}
	return args
}
