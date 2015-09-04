package sorting

import "github.com/xinhuang327/algorithms/common"

func SelectionSort(input []int, delegate SortDelegate) {
	count := len(input)
	for i := 0; i < count-1; i++ {
		min := input[i]
		minIdx := i
		for j := i + 1; j < count; j++ {
			if input[j] < min {
				min = input[j]
				minIdx = j
			}
			delegate.Mark(j)
			delegate.InnerStep(input)
		}
		delegate.Swap(input, i, minIdx)
		common.SwapSlice(input, i, minIdx)

		delegate.Step(input)
	}
}

func BubbleSort(input []int, delegate SortDelegate) {
	count := len(input)
	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if input[j] > input[j+1] {
				delegate.Swap(input, j, j+1)
				common.SwapSlice(input, j, j+1)
			}
			delegate.InnerStep(input)
		}
		delegate.Step(input)
	}
}
