package selecting

func QuickSelect(slice []int, k int) (idx, v int) {
	return quickSelect(slice, 0, len(slice)-1, k)
}

// Returns the kth smallest element in slice[l,r]
// k's range: [1, r-l+1]
func quickSelect(slice []int, l, r int, k int) (idx, v int) {
	s := LomutoPartition(slice, l, r)
	leftSize := s - l // left, pivot, right
	if leftSize == k-1 {
		return s, slice[s]
	}
	if leftSize > k-1 {
		return quickSelect(slice, l, s-1, k)
	} else {
		return quickSelect(slice, s+1, r, k-(leftSize+1)) // left and pivot already smaller
	}
}
