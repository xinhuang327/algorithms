package sorting

import (
	. "github.com/xinhuang327/algorithms/common"
	"github.com/xinhuang327/algorithms/selecting"
)

func QuickSort(slice []int, delegate Delegate) {
	quickSort(slice, 0, len(slice)-1, delegate)
}

func quickSort(slice []int, l, r int, delegate Delegate) {
	if l < r {
		s := selecting.LomutoPartition(slice, l, r, delegate)
		quickSort(slice, l, s-1, delegate)
		quickSort(slice, s+1, r, delegate)
		delegate.Step(slice)
	}
}
