package selecting

import (
	"testing"

	"github.com/xinhuang327/algorithms/common"
	"github.com/xinhuang327/algorithms/sorting"
)

func Test_QuickSelect(t *testing.T) {
	slice := common.GenerateRandomNumbers(50, 1, 100)
	// slice := common.GenerateNumbers(50)
	k := 31

	sorted := sorting.ExecuteSort(sorting.ShellSort, slice, nil)
	kth := sorted[k-1]

	t.Logf("finding %dth\n", k)
	i, n := QuickSelect(slice, k)
	t.Log("slice after", slice)
	t.Log("i", i, "n", n)
	if n != kth {
		t.Error("Wrong selection")
	}
}
