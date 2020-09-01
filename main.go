package main

import (
	"fmt"
)

func main() {
	testSelect("in calling WrongSelect", WrongSelect)
	testSelect("in calling RightSelect", RightSelect)
}

func testSelect(msg string, f func(input1 chan int, input2 chan int) int) {
	var count1, count2 int
	times := 1000
	sourceA := prepareSourceData(1, times)
	sourceB := prepareSourceData(2, times)
	for i := 0; i < times; i++ {
		out := f(sourceA, sourceB)
		if out == 1 {
			count1++
		}
		if out == 2 {
			count2++
		}
	}
	fmt.Println(msg, times, "times, there are ", count1, "'s use input1,", count2, "'s use input2.")
}

func prepareSourceData(data int, times int) chan int {
	input1 := make(chan int, times)
	for i := 0; i < times; i++ {
		input1 <- data
	}
	return input1
}

func WrongSelect(input1, input2 chan int) int {
	select {
	case out := <-input1:
		return out
	case out := <-input2:
		return out
	}
}

func RightSelect(input1 chan int, input2 chan int) int {
	select {
	case out := <-input1:
		return out
	default:
		return <-input2
	}
}
