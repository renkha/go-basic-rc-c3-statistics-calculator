package main

import "fmt"

func main() {
	var n = []int{-6, -4, 0, 1, 2, 2, 3, 10, 12, 44, 52, 72, 98}
	var sc statisticsCalculator
	sc = Array{n}

	fmt.Println(sc.(Array).number)
	fmt.Println(sc.min())
	fmt.Println(sc.max())
	fmt.Println(sc.mean())
	fmt.Println(sc.median())
}

type statisticsCalculator interface {
	min() (string, error)
	max() (string, error)
	mean() (string, error)
	median() (string, error)
}

type Array struct {
	number []int
}

func (a Array) min() (string, error) {
	return "", nil
}

func (a Array) max() (string, error) {
	return "", nil
}

func (a Array) mean() (string, error) {
	return "", nil
}

func (a Array) median() (string, error) {
	return "", nil
}
