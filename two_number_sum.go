package main

import "sort"

// dataset: [6, 87, 1, 4, 8, 7, 9, -1, 19] // unsorted and unique
// target: 8
// quize:
// Is there a pair of two number in dataset sum to target 8?
// answer:
// Yes: 9 + (-1) = 8; [9, -1]

// O(n2) time | O(1) space
func sum2NumberA(dataSet []int, target int) []int {
	size := len(dataSet)
	for i, v := range dataSet {
		for j := i + 1; j < size; j++ {
			b := dataSet[j]
			if v+b == target {
				return []int{v, b}
			}
		}
	}
	return nil
}

// O(n) time | O(n) space
func sum2NumberB(dataSet []int, target int) []int {
	numbers := make(map[int]struct{})
	for _, v := range dataSet {
		if _, existed := numbers[target-v]; existed {
			return []int{v, target - v}
		} else {
			numbers[v] = struct{}{}
		}
	}
	return nil
}

// O(nlogn) time | O(1) space
func sum2NumberC(dataSet []int, target int) []int {
	sort.Ints(dataSet)
	left := 0
	right := len(dataSet) - 1
	for left < right {
		a := dataSet[left]
		b := dataSet[right]
		if a+b == target {
			return []int{a, b}
		} else if a+b > target {
			right--
		} else {
			left++
		}
	}
	return nil
}
