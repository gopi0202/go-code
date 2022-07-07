package utility

import (
	"fmt"
	"math"
	"sort"
)

func MinMaxNumbers(numArray []int, quantifier int, isMin bool) []int {
	if isMin {
		sort.Ints(numArray)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(numArray)))
	}
	output := []int{}
	for i := 0; i < quantifier; i++ {
		output = append(output, numArray[i])
	}
	return output
}

func Average(numArray []int) int {
	total := 0
	for _, num := range numArray {
		total = total + num
	}
	average := total / len(numArray)
	return average
}

func Median(numArray []int) int {
	sort.Ints(numArray)
	midNum := len(numArray) / 2
	if midNum/2 == 0 {
		return (numArray[midNum-1] + numArray[midNum]) / 2
	}
	return numArray[midNum]
}

func Percentaile(numArray []int, quantifier int) int {
	sort.Ints(numArray)
	fmt.Println(quantifier, numArray)
	percentaile := (float64(quantifier) / 100) * float64((len(numArray) - 1))
	roundKey := math.Ceil(float64(percentaile))
	return numArray[int64(roundKey)]
}
