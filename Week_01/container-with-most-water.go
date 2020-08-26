package main

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	 maxArea(height)
	 maxArea1(height)
}

func maxArea1(height []int) int {
	//双循环优化
	var (
		i,j = 0,len(height) - 1
		max = 0
	)
	for j > i {
		if height[j] < height[i] {
			area := height[j] * (j - i)
			if max < area {
				max = area
			}
			j--
		} else {
			area := height[i] * (j - i)
			if max < area {
				max = area
			}
			i++
		}
	}

	return max
}

func maxArea(height []int) int {
	//双循环法
	max := 0
	l := len(height)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			min := 0

			if height[i] < height[j] {
				min = height[i]
			} else {
				min = height[j]
			}

			tmp := min * (j - i)
			if max < tmp {
				max = tmp
			}
		}
	}

	return max
}
