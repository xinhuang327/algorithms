package sorting

import (
	"fmt"

	. "github.com/xinhuang327/algorithms/common"
)

type SortFunc func(input []int, delegate Delegate)

func ExecuteSort(sort SortFunc, input []int, delegate Delegate) []int {
	del := delegate
	if del == nil {
		del = &NullDelegate{}
	}
	inputClone := append([]int{}, input...)

	del.Init(inputClone)
	sort(inputClone, del)
	del.Finish(inputClone)

	finished := inputClone

	isSorted, errIdx := VerifySorted(finished)
	if !isSorted {
		fmt.Println("Not Sorted, Error Index: ", errIdx)
	} else {
		fmt.Println("Sorted.")
	}

	return inputClone
}

func VerifySorted(input []int) (isSorted bool, errIdx int) {
	count := len(input)
	for i := 0; i < count-1; i++ {
		if input[i] > input[i+1] {
			return false, i
		}
	}
	return true, -1
}
