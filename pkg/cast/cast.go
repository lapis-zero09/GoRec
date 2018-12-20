package cast

import (
	"log"
	"strconv"
)

func IntSliceToFloatSlice(i []int) []float64 {
	f := make([]float64, cap(i))
	for j, val := range i {
		f[j] = float64(val)
	}
	return f
}

func StrSliceToIntslice(s []string) []int {
	intS := make([]int, cap(s))
	for i, val := range s {
		val, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		intS[i] = val
	}
	return intS
}
