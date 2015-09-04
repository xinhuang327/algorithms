package selecting

import "github.com/xinhuang327/algorithms/common"

// Returns partition of slice[l,r] and pivot position s, pivot is 1st element of original slice
// for i < s, slice[i] < pivot, i == s, slice[i] == pivot, i > s, slice[i] >= pivot
func LomutoPartition(slice []int, l, r int) (s int) {
	p := slice[l]
	s = l
	for i := l + 1; i <= r; i++ {
		if slice[i] < p {
			s++
			common.SwapSlice(slice, i, s)
		}
	}
	common.SwapSlice(slice, s, l)
	return
}
