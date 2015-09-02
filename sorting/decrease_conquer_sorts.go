package sorting

func InsertionSort(input []int, delegate SortDelegate) {
	count := len(input)
	for i := 1; i < count; i++ {
		v := input[i]
		j := i - 1
		for j >= 0 && input[j] > v {
			input[j+1] = input[j]
			delegate.Set(j+1, input[j])

			j--
			delegate.InnerStep(input)
		}
		input[j+1] = v
		delegate.Set(j+1, v)

		delegate.Step(input)
	}
}

var ShellSortGaps = []int{701, 301, 132, 57, 23, 10, 4, 1}

// var ShellSortGaps = []int{1}

func ShellSort(input []int, delegate SortDelegate) {
	count := len(input)
	for _, gap := range ShellSortGaps {
		for i := gap; i < count; i++ {
			v := input[i]
			j := i - gap
			for j >= 0 && input[j] > v {
				input[j+gap] = input[j]
				delegate.Set(j+gap, input[j])

				j -= gap
				delegate.InnerStep(input)
			}
			input[j+gap] = v
			delegate.Set(j+gap, v)

			delegate.Step(input)
		}
	}
}
