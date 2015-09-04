package main

import (
	"github.com/xinhuang327/algorithms/common"
	. "github.com/xinhuang327/algorithms/sorting"
)

func main() {
	nums := common.GenerateRandomNumbers(1000000, 10, 100000)
	delegate := &common.PrintDelegate{}
	// ExecuteSort(BubbleSort, nums, delegate)
	// ExecuteSort(SelectionSort, nums, delegate)
	// ExecuteSort(InsertionSort, nums, delegate)
	// ExecuteSort(ShellSort, nums, delegate)
	ExecuteSort(QuickSort, nums, delegate)
	ExecuteSort(MergeSort, nums, delegate)
}
