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
var winW = 1600
var winH = 1000
var numBars = 100
var valNum = 10000

var defaultBgBrush = gxui.CreateBrush(gxui.White)
var markBgBrush = gxui.CreateBrush(gxui.Red)
var setBgBrush = gxui.CreateBrush(gxui.Green)
var swapABgBrush = gxui.CreateBrush(gxui.Blue)
var swapBBgBrush = gxui.CreateBrush(gxui.Blue50)
var bars []gxui.Button

func createBar() gxui.Button {
	child := theme.CreateButton()
	child.SetBackgroundBrush(defaultBgBrush)
	w := winW / numBars
	if w < 1 {
		w = 1
	}
	child.SetBorderPen(gxui.TransparentPen)
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
	layout.SetBackgroundBrush(gxui.CreateBrush(gxui.Black))
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
		// sorting.ExecuteSort(sorting.ShellSort, nums, delegate)
		// sorting.ExecuteSort(sorting.MergeSort, nums, delegate)
		sorting.ExecuteSort(sorting.QuickSort, nums, delegate)
		// fmt.Println(result)
	}()
}

func main() {
	gl.StartDriver(appMain)
}

type GUIDelegate struct {
	lastMarkIdx  int
	lastSetIdx   int
	lastSwapIdxA int
	lastSwapIdxB int
}

func NewGUIDelegate() *GUIDelegate {
	return &GUIDelegate{
		-1,
		-1,
		-1,
		-1,
	}
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

		if p.lastSetIdx >= 0 {
			bars[p.lastSetIdx].SetBackgroundBrush(defaultBgBrush)
		}

		bars[idx].SetBackgroundBrush(setBgBrush)
		p.lastSetIdx = idx
	})
	wait()
}
func (p *GUIDelegate) Swap(slice []int, i, j int) {
	iv := slice[i]
	jv := slice[j]
	theDriver.Call(func() {
		setBarHeight(bars[i], jv)
		setBarHeight(bars[j], iv)

		if p.lastSwapIdxA >= 0 {
			bars[p.lastSwapIdxA].SetBackgroundBrush(defaultBgBrush)
		}
		if p.lastSwapIdxB >= 0 {
			bars[p.lastSwapIdxB].SetBackgroundBrush(defaultBgBrush)
		}

		bars[i].SetBackgroundBrush(swapABgBrush)
		bars[j].SetBackgroundBrush(swapBBgBrush)
		p.lastSwapIdxA = i
		p.lastSwapIdxB = j
	})
	wait()
}
func (p *GUIDelegate) Mark(idx int) {
	theDriver.Call(func() {
		if p.lastMarkIdx >= 0 {
			bars[p.lastMarkIdx].SetBackgroundBrush(defaultBgBrush)
		}
		bars[idx].SetBackgroundBrush(markBgBrush)
		p.lastMarkIdx = idx
	})
	wait()
}

func wait() {
	ms := time.Duration(5000 / numBars)
	<-time.After(ms * time.Millisecond)
}
