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
	i := make([]int, cap(s))
	for j, val := range s {
		val, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		i[j] = val
	}
	return i
}
