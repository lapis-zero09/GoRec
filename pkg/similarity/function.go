package similarity

import (
	"math"

	"github.com/lapis-zero09/GoRec/pkg/cast"
)

func Cosine(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(tmpVec1)) * math.Sqrt(SumSquad(tmpVec2))
	return (numerator / deliminator), nil
}

func AdjustedCosine(vec1, vec2 []int, mean ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	adjustedM, err := SliceSub(tmpVec1, mean[0])
	adjustedN, err := SliceSub(tmpVec2, mean[0])
	if err != nil {
		return 0.0, err
	}
	numerator, err := Dot(adjustedM, adjustedN)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(adjustedM)) * math.Sqrt(SumSquad(adjustedN))
	return (numerator / deliminator), nil
}

func Jaccard(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := (Sum(tmpVec1) + Sum(tmpVec2)) - numerator
	return (numerator / deliminator), nil
}

func Dice(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := Sum(tmpVec1) + Sum(tmpVec2)
	return (2 * numerator / deliminator), nil
}

func Simpson(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	sum := Sum(tmpVec1)
	if tmpSum := Sum(tmpVec2); sum > tmpSum {
		return (numerator / tmpSum), nil
	}
	return (numerator / sum), nil
}

func Pearson(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := cast.IntSliceToFloatSlice(vec1)
	tmpVec2 := cast.IntSliceToFloatSlice(vec2)
	uDiff := SubMean(tmpVec1)
	vDiff := SubMean(tmpVec2)
	numerator, err := Dot(uDiff, vDiff)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(uDiff)) * math.Sqrt(SumSquad(vDiff))
	return (numerator / deliminator), nil
}
