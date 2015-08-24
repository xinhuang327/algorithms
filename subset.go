package main

import (
	"fmt"

	"github.com/xinhuang327/algorithms/common"
)

func main() {
	n := 5
	elems := common.GenerateNumbers(n)
	// fmt.Println(elems)

	flags := common.GenerateFlagNumbers(n) // 1,2,4,8....
	count := common.Power(2, n)

	allSubsets := [][]int{}
	for i := 0; i < count; i++ {
		subset := []int{}
		for k, flag := range flags {
			if i&flag != 0 {
				subset = append(subset, elems[k])
			}
		}
		allSubsets = append(allSubsets, subset)
	}

	for _, subset := range allSubsets {
		fmt.Println(subset)
	}
}
