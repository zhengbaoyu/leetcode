package main

import (
	"fmt"
	"strings"
)

/*
	给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。
	如果不存在，则返回 -1。

	思路：核心点遍历给定字符串字符，判断以当前字符开头字符串是否等于目标字符串
*/

func main() {
	haystack := "ooloooooo"
	needle := "l"
	value := strStr3(haystack, needle)
	fmt.Println(value)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	// i不需要到len-1
	for i = 0; i < len(haystack); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				fmt.Println("====", haystack[i+j], needle[j])
				break
			}
		}
		// 判断字符串长度是否相等
		if len(needle) == j {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	length := len(needle) // 提取，循环中不计算长度
	if length == 0 {      // 加速，可以去掉
		return 0
	}
	length2 := len(haystack) // 提取，循环中不计算长度
	for k := range haystack {
		if k+length <= length2 {
			if haystack[k:k+length] == needle {
				return k
			}
		} else { // 加速，可以去掉
			return -1
		}
	}
	return -1
}

func strStr3(haystack string, needle string) int {
	haystack = "_000_"
	needle = "1"
	return strings.Index(haystack, needle)
}

// 最优解题方案
func strStr4(haystack string, needle string) int {

	l1 := len(haystack)
	l2 := len(needle)

	if l2 == 0 {
		return 0
	}
	if l1 == 0 || l1 < l2 {
		return -1
	}

	for i := 0; i <= l1 - l2; i++ {
		fmt.Println(haystack[i:i+l2])
		if haystack[i : i + l2] == needle {
			return i
		}
	}
	return -1
}
