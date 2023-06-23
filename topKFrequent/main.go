package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 3}, 3))
}

func topKFrequent(nums []int, k int) []int {
	top := make(map[int]int)
	for _, num := range nums {
		top[num]++
	}

	// fmt.Println(top)

	getK := make([]int, k)
	for i := 0; i < k; i++ {
		for num, count := range top {
			if count > top[getK[i]] {
				getK[i] = num
			}
		}
		delete(top, getK[i])
	}

	// fmt.Println(getK)

	return getK
}

// Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
