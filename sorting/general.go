package sorting

import (
	"fmt"
	"time"
)

type SortDelegate interface {
	Init(input []int)
	Step(current []int)
	InnerStep(current []int)
	Finish(finished []int)
}

type SortFunc func(input []int, delegate SortDelegate)

type NullSortDelegate struct {
}

func (p *NullSortDelegate) Init(input []int) {
}
func (p *NullSortDelegate) Step(current []int) {
}
func (p *NullSortDelegate) InnerStep(current []int) {
}
func (p *NullSortDelegate) Finish(finished []int) {
}

type PrintSortDelegate struct {
	StepCount         int
	InnerStepCount    int
	ReportStep        bool
	ReportInputOutput bool
	StartTime         time.Time
}

func (p *PrintSortDelegate) Init(input []int) {
	p.StepCount = 0
	p.InnerStepCount = 0
	if p.ReportInputOutput {
		fmt.Println("Input: ", input)
	}
	p.StartTime = time.Now()
}
func (p *PrintSortDelegate) Step(current []int) {
	if p.ReportStep {
		fmt.Println("Step: ", current)
	}
	p.StepCount += 1
}
func (p *PrintSortDelegate) InnerStep(current []int) {
	p.InnerStepCount += 1
}
func (p *PrintSortDelegate) Finish(finished []int) {
	fmt.Println("Steps: ", p.StepCount, ", Input Length:", len(finished), ", Inner Steps", p.InnerStepCount)
	isSorted, errIdx := VerifySorted(finished)
	if !isSorted {
		fmt.Println("Not Sorted, Error Index: ", errIdx)
		if p.ReportInputOutput {
			fmt.Println(finished)
		}
	} else {
		if p.ReportInputOutput {
			fmt.Println("Finish: ", finished)
		}
	}
	elapsed := time.Since(p.StartTime)
	fmt.Println("Elapsed: ", elapsed)
}

func ExecuteSort(sort SortFunc, input []int, delegate SortDelegate) []int {
	del := delegate
	if del == nil {
		del = &NullSortDelegate{}
	}
	inputClone := append([]int{}, input...)

	del.Init(inputClone)
	sort(inputClone, del)
	del.Finish(inputClone)
	return inputClone
}

func VerifySorted(input []int) (isSorted bool, errIdx int) {
	count := len(input)
	for i := 0; i < count-1; i++ {
		if input[i] > input[i+1] {
			return false, i
		}
	}
	return true, -1
}
