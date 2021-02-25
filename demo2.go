package main

import "fmt"

/*
	给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i,ai) 。
	在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
	找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
	输入：[1,8,6,2,5,4,8,3,7]
	输出：49
*/
func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	areaV := area(nums)
	fmt.Println(areaV)
}

func area(num []int) int {
	max := 0
	for k, v := range num {
		for i, j := range num {
			w := k - i
			if w < 0 {
				w = -1 * w
			}
			h := v
			if v > j {
				h = j
			}
			
			if max < w*h {
				max = w * h
			}
		}
	}
	return max
}
