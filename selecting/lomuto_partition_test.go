package selecting

import (
	"testing"

	"github.com/xinhuang327/algorithms/common"
)

func Test_LomutoPartition(t *testing.T) {
	slice := common.GenerateRandomNumbers(100, 1, 100)
	l := 13
	r := len(slice) - 7
	pivot := slice[l]
	t.Log("slice", slice)
	s := LomutoPartition(slice, l, r)
	t.Log("partitioned", slice)
	for i := l; i <= r; i++ {
		a := slice[i]
		if i < s && a >= pivot || i == s && a != pivot || i > s && a < pivot {
			t.Error("Wrong partition:", i, a)
		}
	}
}
