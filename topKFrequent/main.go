package main

import (
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 3))
	//fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3, 3}, 3))
	//fmt.Println(topKFrequent3([]int{1, 1, 1, 2, 2, 3, 3}, 3))
	//fmt.Println(topKFrequent4([]int{1, 1, 1, 2, 2, 3, 3}, 3))
}

// solution 1
// O(^n2)
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	fmt.Println(freq)

	getK := make([]int, k)
	for i := 0; i < k; i++ {
		for num, count := range freq {
			if count > freq[getK[i]] {
				getK[i] = num
			}
			fmt.Println(num, count, getK)
		}
		delete(freq, getK[i])
	}

	// fmt.Println(getK)

	return getK
}

////solution 2
////O(n)
//func topKFrequent2(nums []int, k int) []int {
//	freq := make(map[int]int)
//	for _, num := range nums {
//		freq[num]++
//	}
//
//	getK := make([]int, k)
//	for num, count := range freq {
//		if count > freq[getK[0]] {
//			getK[0] = num
//			for i := 0; i < k-1; i++ {
//				if freq[getK[i]] > freq[getK[i+1]] {
//					getK[i], getK[i+1] = getK[i+1], getK[i]
//				}
//			}
//		}
//	}
//
//	return getK
//}
//
//// solution 3
//// O(nlogk)
//func topKFrequent3(nums []int, k int) []int {
//	freqMap := make(map[int]int)
//
//	// Count the frequency of each number
//	for _, num := range nums {
//		freqMap[num]++
//	}
//
//	// Create a list of unique numbers
//	uniqueNums := make([]int, 0, len(freqMap))
//	for num := range freqMap {
//		uniqueNums = append(uniqueNums, num)
//	}
//
//	// Sort the unique numbers based on their frequency
//	sort.Slice(uniqueNums, func(i, j int) bool {
//		return freqMap[uniqueNums[i]] > freqMap[uniqueNums[j]]
//	})
//
//	// Return the top K frequent numbers
//	return uniqueNums[:k]
//}
//
//// solution 4
//// O(nlogn)
//
//type Node struct {
//	val  int
//	freq int
//}
//
//type PriorityQueue []*Node
//
//func (pq PriorityQueue) Len() int {
//	return len(pq)
//}
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].freq < pq[j].freq
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	*pq = append(*pq, x.(*Node))
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	x := old[n-1]
//	*pq = old[:n-1]
//	return x
//}
//
//func topKFrequent4(nums []int, k int) []int {
//	// Create a map to store the frequency of each element.
//	occurrence := make(map[int]int)
//	for _, v := range nums {
//		occurrence[v]++
//	}
//
//	// Create a priority queue to store the top k frequent elements.
//	pq := make(PriorityQueue, 0, k)
//	for v, freq := range occurrence {
//		if len(pq) < k {
//			heap.Push(&pq, &Node{val: v, freq: freq})
//		} else if freq > pq[0].freq {
//			heap.Pop(&pq)
//			heap.Push(&pq, &Node{val: v, freq: freq})
//		}
//	}
//
//	// Get the elements from the priority queue.
//	res := make([]int, 0, k)
//	for pq.Len() > 0 {
//		node := heap.Pop(&pq).(*Node)
//		res = append(res, node.val)
//	}
//
//	return res
//}
