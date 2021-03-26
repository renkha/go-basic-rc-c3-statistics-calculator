package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n = []int{-6, -4, 0, 1, 2, 2, 3, 10, 12, 44, 52, 72, 98}
	var sc statisticsCalculator
	sc = Array{n}

	fmt.Println(sc.(Array).number)
	fmt.Println(sc.min())
	fmt.Println(sc.max())
	fmt.Println(sc.median())
	fmt.Println(sc.mean())
}

type statisticsCalculator interface {
	min() (string, error)
	max() (string, error)
	median() (string, error)
	mean() (string, error)
}

type Array struct {
	number []int
}

func (a Array) min() (string, error) {
	min := a.number[0]
	for i := 0; i < len(a.number); i++ {
		if min >= a.number[i] {
			min = a.number[i]
		}
	}
	return "Min: " + strconv.Itoa(min), nil
}

func (a Array) max() (string, error) {
	max := a.number[0]
	for i := 0; i < len(a.number); i++ {
		if max <= a.number[i] {
			max = a.number[i]
		}
	}
	return "Max: " + strconv.Itoa(max), nil
}

func (a Array) median() (string, error) {
	return "", nil
}

func (a Array) mean() (string, error) {
	return "", nil
}
