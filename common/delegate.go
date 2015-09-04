package common

import (
	"fmt"
	"time"
)

type Delegate interface {
	Init(input []int)
	Step(current []int)
	InnerStep(current []int)
	Finish(finished []int)
	Set(idx int, val int)
	Swap(slice []int, i, j int)
	Mark(idx int)
}

type NullDelegate struct {
}

func (p *NullDelegate) Init(input []int) {
}
func (p *NullDelegate) Step(current []int) {
}
func (p *NullDelegate) InnerStep(current []int) {
}
func (p *NullDelegate) Finish(finished []int) {
}
func (p *NullDelegate) Set(idx int, val int) {
}
func (p *NullDelegate) Swap(slice []int, i, j int) {
}
func (p *NullDelegate) Mark(idx int) {
}

type PrintDelegate struct {
	StepCount      int
	InnerStepCount int
	ReportStep     bool
	StartTime      time.Time
}

func (p *PrintDelegate) Init(input []int) {
	p.StepCount = 0
	p.InnerStepCount = 0
	p.StartTime = time.Now()
}
func (p *PrintDelegate) Step(current []int) {
	if p.ReportStep {
		fmt.Println("Step: ", current)
	}
	p.StepCount += 1
}
func (p *PrintDelegate) InnerStep(current []int) {
	p.InnerStepCount += 1
}
func (p *PrintDelegate) Finish(finished []int) {
	fmt.Println("Steps: ", p.StepCount, ", Input Length:", len(finished), ", Inner Steps", p.InnerStepCount)
	elapsed := time.Since(p.StartTime)
	fmt.Println("Elapsed: ", elapsed)
}

func (p *PrintDelegate) Set(idx int, val int) {
}
func (p *PrintDelegate) Swap(slice []int, i, j int) {
}
func (p *PrintDelegate) Mark(idx int) {
}
