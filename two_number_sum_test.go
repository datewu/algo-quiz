package main

import "testing"

func TestTwoNumberSum(t *testing.T) {
	dataset := []int{6, 87, 1, 4, 8, 7, 9, -1, 19, -1876, 1998} // unsorted and unique
	target := 8

	existed := func(result []int) {
		if len(result) != 2 {
			t.Error("expect existed true got false")
		}
	}

	notExisted := func(result []int) {
		if result != nil {
			t.Error("expect existed false got true")
		}
	}
	existed(sum2NumberA(dataset, target))
	existed(sum2NumberB(dataset, target))
	existed(sum2NumberC(dataset, target))

	existed(sum2NumberA(dataset, target+10))
	existed(sum2NumberB(dataset, target+10))
	existed(sum2NumberC(dataset, target+10))

	notExisted(sum2NumberA(dataset, target+30))
	notExisted(sum2NumberB(dataset, target+30))
	notExisted(sum2NumberC(dataset, target+30))
}
