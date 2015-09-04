package sorting

import (
	"math"
)

func MergeSort(slice []int, delegate SortDelegate) {
	mergeSort(slice, delegate, 0)
}

// l: start index on original slice, only for visualization
func mergeSort(slice []int, delegate SortDelegate, l int) {
	if len(slice) > 1 {
		mid := int(math.Floor(float64(len(slice)) / 2.0))
		a := append([]int{}, slice[0:mid]...)
		b := append([]int{}, slice[mid:]...)
		mergeSort(a, delegate, l)
		mergeSort(b, delegate, l+mid)
		merge(a, b, slice, delegate, l)
	}
}

func merge(a, b, dst []int, delegate SortDelegate, l int) {
	i := 0
	j := 0
	k := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			dst[k] = a[i]
			delegate.Set(l+k, a[i]) // add l for original index to visualize
			i++
		} else {
			dst[k] = b[j]
			delegate.Set(l+k, b[j])
			j++
		}
		k++
	}
	if i == len(a) {
		copy(dst[k:], b[j:])
		for j < len(b) {
			delegate.Set(l+k, b[j])
			j++
			k++
		}
	} else {
		copy(dst[k:], a[i:])
		for i < len(a) {
			delegate.Set(l+k, a[i])
			i++
			k++
		}
	}
}
