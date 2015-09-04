package main

import (
	"fmt"

	"github.com/xinhuang327/algorithms/common"
)

func main() {
	perms := LexicographicPermute(4)
	// for _, perm := range perms {
	// 	fmt.Println(perm)
	// }
	fmt.Println("len: ", len(perms))
}

func LexicographicPermute(n int) [][]int {

	var results [][]int
	perm := common.GenerateNumbers(n)

	results = append(results, perm)

	findIJ := func(perm []int) (i, j int) {
		i = -1
		j = -1
		for k := 0; k < len(perm)-1; k++ {
			if perm[k] < perm[k+1] {
				i = k
			}
			if i != -1 && perm[k+1] > perm[i] {
				j = k + 1
			}
		}
		return
	}

	i, j := findIJ(perm)
	fmt.Println(perm, i, j)
	for i != -1 {
		perm = append([]int(nil), perm...)
		common.SwapElement(perm, i, j)
		common.ReverseSliceRange(perm, i+1, n-1)
		results = append(results, perm)
		i, j = findIJ(perm)
		fmt.Println(perm, i, j)
	}

	return results
}
