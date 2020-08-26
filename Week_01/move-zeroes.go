package main

func main() {
	nums := []int{1, 1, 2, 1, 3}

	moveZeroes1(nums)
	moveZeroes2(nums)
	moveZeroes3(nums)
	moveZeroes4(nums)

}

func moveZeroes4(nums []int) {
	//go语言中切割数组

	l := len(nums)
	i := 0

	for i < l {
		if nums[i] == 0 {
			//把0从数组中删掉
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			l--
		} else {
			i++
		}
	}
}

func moveZeroes3(nums []int) {
	//快慢指针

	//j为慢指针，永远指向当前第一个0，用来和 i指向的非0交换
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			//备注，可将后面的nums[j] 替换为0，因为j永远指向为当前第一个0
			//nums[i], nums[j] = 0, nums[i]
			//此时如果nums数组中没有0也不用担心，因为j永远等于i,nums[j]会覆盖 0 的赋值
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

}

func moveZeroes2(nums []int) {
	//双指针法
	i := 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	//双数组暴力解法

	var tmp []int

	for _, v := range nums {
		if v != 0 {
			tmp = append(tmp, v)
		}
	}

	zeroNum := len(nums) - len(tmp)

	for i := 0; i < zeroNum; i++ {
		tmp = append(tmp, 0)
	}

	nums = tmp
}
