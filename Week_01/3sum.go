package main

import "sort"

func main() {
	nums := []int{3,0,-2,-1,1,2}
	threeSum3(nums)
}

func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum3(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//枚举first
	for f1 := 0; f1 < l; f1++ {
		//去重
		if f1 > 0 && nums[f1] == nums[f1-1] {
			continue
		}

		//初始化f3
		f3 := l - 1

		//获取目标值
		target := -1 * nums[f1]

		//如果最左边第一个>0 则和不可能为0
		if target < 0 {
			break
		}

		//遍历
		for f2 := f1 + 1; f2 < l; f2++ {
			//去重
			if f2 > f1 + 1 && nums[f2] == nums[f2-1] {
				continue
			}

			//移动f3,一直到跟f2 相遇或 和 <= 目标值
			for f2 < f3 && nums[f2]+nums[f3] > target {
				f3--
			}

			//相等则退出改循环, f1右移
			if f2 == f3 {
				break
			}

			if nums[f2]+nums[f3] == target {
				ans = append(ans, []int{nums[f1], nums[f2], nums[f3]})
			}
		}

	}

	return ans
}

func threeSum(nums []int) [][]int {
	//暴力解法
	var m [][]int
	l := len(nums)
	sort.Ints(nums)
	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					m = append(m, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return m
}
