package similarity

import "fmt"

func Mean(arr []float64) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += val
	}
	return sum / float64(len(arr))
}

func SubMean(arr []float64) []float64 {
	mean := Mean(arr)
	diff_arr := make([]float64, len(arr))
	for i, val := range arr {
		diff_arr[i] = val - mean
	}
	return diff_arr
}

func Sum(arr []float64) float64 {
	var res float64
	for _, val := range arr {
		res += val
	}
	return res
}

func SumSquad(arr []float64) float64 {
	var res float64
	for _, val := range arr {
		res += (val * val)
	}
	return res
}

func Dot(u, v []float64) (float64, error) {
	if len(u) != len(v) {
		return 0.0, fmt.Errorf("Array size is different between u and v")
	}

	var res float64
	for i := 0; i < len(u); i++ {
		res += float64(u[i]) * float64(v[i])
	}
	return res, nil
}

func SliceSub(u, v []float64) ([]float64, error) {
	if len(u) != len(v) {
		return nil, fmt.Errorf("Array size is different between u and v")
	}

	res := make([]float64, len(u))
	for i := 0; i < len(u); i++ {
		res[i] = float64(u[i]) - float64(v[i])
	}
	return res, nil
}
