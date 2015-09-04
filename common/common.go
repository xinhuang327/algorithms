package common

import (
	"math"
	"math/rand"
	"time"
)

func Power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func GenerateNumbers(count int) []int {
	var nums []int
	for i := 1; i <= count; i++ {
		nums = append(nums, i)
	}
	return nums
}

func GenerateFlagNumbers(count int) []int {
	var nums []int
	for i := 0; i < count; i++ {
		nums = append(nums, Power(2, i))
	}
	return nums
}

func GenerateRandomNumbers(count int, min int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	var nums []int
	for i := 0; i < count; i++ {
		n := min + int(math.Floor(rand.Float64()*float64(max-min)))
		nums = append(nums, n)
	}
	return nums
}

func ReverseSliceRange(slice []int, a, b int) {
	for i := a; i <= (a+b)/2; i++ {
		j := b - (i - a)
		SwapSlice(slice, i, j)
	}
}

func SwapSlice(slice []int, i, j int) {
	if i == j {
		return
	}
	t := slice[i]
	slice[i] = slice[j]
	slice[j] = t
}
