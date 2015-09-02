package main

import (
	"fmt"
	"time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
	"github.com/xinhuang327/algorithms/common"
	"github.com/xinhuang327/algorithms/sorting"
)

var theDriver gxui.Driver
var theme gxui.Theme
var winW = 1500
var winH = 600
var numBars = 1400
var valNum = 10000

var bars []gxui.Button

func createBar() gxui.Button {
	child := theme.CreateButton()
	child.SetBackgroundBrush(gxui.CreateBrush(gxui.White))
	w := winW / numBars
	if w < 1 {
		w = 1
		child.SetBorderPen(gxui.TransparentPen)
	}
	child.SetPadding(math.Spacing{L: w, T: w})
	child.SetMargin(math.ZeroSpacing)
	return child
}

func setBarHeight(b gxui.Button, n int) {
	// fmt.Println("setBarHeight")
	h := float64(n) / float64(valNum)
	spacing := b.Padding()
	spacing.T = int(h * float64(winH))
	b.SetPadding(spacing)
	b.Relayout()
}

func appMain(driver gxui.Driver) {
	theDriver = driver
	theme = flags.CreateTheme(driver)

	window := theme.CreateWindow(winW, winH, "Window")
	window.OnClose(driver.Terminate)
	window.SetScale(flags.DefaultScaleFactor)

	layout := theme.CreateLinearLayout()
	layout.SetBackgroundBrush(gxui.CreateBrush(gxui.Blue40))
	layout.SetDirection(gxui.LeftToRight)
	layout.SetVerticalAlignment(gxui.AlignBottom)

	layout.SetSizeMode(gxui.Fill)

	nums := common.GenerateRandomNumbers(numBars, 0, valNum)
	for _, n := range nums {
		child := createBar()
		setBarHeight(child, n)
		layout.AddChild(child)
		bars = append(bars, child)
	}

	window.AddChild(layout)

	delegate := &GUIDelegate{}

	go func() {
		<-time.After(1 * time.Second)
		fmt.Println("ExecuteSort...")
		// sorting.ExecuteSort(sorting.InsertionSort, nums, delegate)
		// sorting.ExecuteSort(sorting.BubbleSort, nums, delegate)
		// sorting.ExecuteSort(sorting.SelectionSort, nums, delegate)
		sorting.ExecuteSort(sorting.ShellSort, nums, delegate)
	}()
}

func main() {
	gl.StartDriver(appMain)
}

type GUIDelegate struct {
}

func (p *GUIDelegate) Init(input []int) {
}
func (p *GUIDelegate) Step(current []int) {
}
func (p *GUIDelegate) InnerStep(current []int) {
}
func (p *GUIDelegate) Finish(finished []int) {
}
func (p *GUIDelegate) Set(idx int, val int) {
	theDriver.Call(func() {
		setBarHeight(bars[idx], val)
	})
	wait()
}
func (p *GUIDelegate) Swap(slice []int, i, j int) {
	iv := slice[i]
	jv := slice[j]
	theDriver.Call(func() {
		setBarHeight(bars[i], jv)
		setBarHeight(bars[j], iv)
	})
	wait()
}

func wait() {
	<-time.After(1 * time.Millisecond)
}
