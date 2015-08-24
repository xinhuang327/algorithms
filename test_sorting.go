package main

import (
	"github.com/xinhuang327/algorithms/common"
	. "github.com/xinhuang327/algorithms/sorting"
)

func main() {
	nums := common.GenerateRandomNumbers(100000, 10, 100000)
	delegate := &PrintSortDelegate{}
	// ExecuteSort(BubbleSort, nums, delegate)
	// ExecuteSort(SelectionSort, nums, delegate)
	ExecuteSort(InsertionSort, nums, delegate)
	ExecuteSort(ShellSort, nums, delegate)
}
