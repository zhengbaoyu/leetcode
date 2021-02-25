package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

func main() {
	nums := []int{3, 2, 4}
	target := 6
	value := twoSum2(nums, target)
	fmt.Println(value)
}

//标准解题法
func twoSum(nums []int, target int) []int {
	kMap := make(map[int]int) //map[value]key

	for k, v := range nums {

		if mapValue, ok := kMap[v]; ok {
			fmt.Println(kMap)
			return []int{mapValue, k}
		}
		kMap[target-v] = k
	}
	return nil
}

// 嵌套循环解题法
func twoSum2(nums []int, target int) []int {
	for k := range nums {
		for i := 0; i < k; i++ {
			if nums[k] + nums[i] == target{
				return []int{k,i}
			}
		}
	}
	return nil
}
