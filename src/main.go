package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"

	"github.com/lapis-zero09/GoRec/src/preprocessing"
)

type Encountered struct {
	UniqueUser []int
	UniqueItem []int
}

func StrsliceToIntslice(string_slice []string) []int {
	int_array := make([]int, cap(string_slice))
	for i, val := range string_slice {
		val, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		}
		int_array[i] = val
	}
	return int_array
}

func IntsliceToFloatslice(int_slice []int) []float64 {
	float_array := make([]float64, cap(int_slice))
	for i, val := range int_slice {
		float_array[i] = float64(val)
	}
	return float_array
}

func ReadFileToData(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file can't be opened from os open")
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'
	var data [][]int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't found EOF")
		}
		data_int := StrsliceToIntslice(record[:3])
		data = append(data, data_int)
	}
	return data, nil
}

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

func TakeCol(data [][]int, colIdx int) ([]int, error) {
	if colIdx < 0 || len(data[0]) < colIdx {
		return nil, fmt.Errorf("column index is invalid value!")
	}
	itemVec := make([]int, len(data))
	for i, val := range data {
		itemVec[i] = val[colIdx]
	}
	return itemVec, nil
}

func Cosine(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(tmpVec1)) * math.Sqrt(SumSquad(tmpVec2))
	return (numerator / deliminator), nil
}

func AdjustedCosine(vec1, vec2 []int, mean ...[]float64) (float64, error) {
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
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
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := (Sum(tmpVec1) + Sum(tmpVec2)) - numerator
	return (numerator / deliminator), nil
}

func Dice(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
	numerator, err := Dot(tmpVec1, tmpVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := Sum(tmpVec1) + Sum(tmpVec2)
	return (2 * numerator / deliminator), nil
}

func Simpson(vec1, vec2 []int, _ ...[]float64) (float64, error) {
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
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
	tmpVec1 := IntsliceToFloatslice(vec1)
	tmpVec2 := IntsliceToFloatslice(vec2)
	uDiff := SubMean(tmpVec1)
	vDiff := SubMean(tmpVec2)
	numerator, err := Dot(uDiff, vDiff)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(uDiff)) * math.Sqrt(SumSquad(vDiff))
	return (numerator / deliminator), nil
}

func MakeSimilarityMatrix(userItemMatrix [][]int, method func([]int, []int, ...[]float64) (float64, error), userFlag bool) [][]float64 {
	var size int
	if userFlag {
		size = len(userItemMatrix)
	} else {
		size = len(userItemMatrix[0])
	}
	simMat := make([][]float64, size)
	for i := 0; i < size; i++ {
		simMat[i] = make([]float64, size)
	}

	var mean []float64
	funcName := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
	if funcName == "github.com/lapis-zero09/GoRec/src.AdjustedCosine" {
		if userFlag {
			mean = make([]float64, len(userItemMatrix[0]))
			for i := 0; i < len(userItemMatrix[0]); i++ {
				tmpVec, err := TakeCol(userItemMatrix, i)
				if err != nil {
					fmt.Println(err)
					break
				}
				mean[i] = Mean(IntsliceToFloatslice(tmpVec))
			}
		} else {
			mean = make([]float64, len(userItemMatrix))
			for i, val := range userItemMatrix {
				mean[i] = Mean(IntsliceToFloatslice(val))
			}
		}
	}
	var err error
	var sim float64
	for i := 0; i < size; i++ {
		for j := size - 1; j > i; j-- {
			if userFlag {
				sim, err = method(userItemMatrix[i], userItemMatrix[j], mean)
			} else {
				iVec, err := TakeCol(userItemMatrix, i)
				jVec, err := TakeCol(userItemMatrix, j)
				if err != nil {
					fmt.Println(err)
					break
				}
				sim, err = method(iVec, jVec, mean)
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			simMat[i][j] = sim
			simMat[j][i] = sim
		}
	}
	return simMat
}

func MostSimilar(encounteredUnique []int, simMat [][]float64, id, similarSize int) map[int]float64 {
	simVector := simMat[id]
	sortedSimVector := make([]float64, len(simVector))
	copy(sortedSimVector, simVector)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedSimVector)))

	ranker := make(map[int]float64, similarSize)
	for _, sortedSim := range sortedSimVector[:similarSize] {
		for i, sim := range simVector {
			if sim == sortedSim {
				ranker[encounteredUnique[i]] = sim
			}
		}
	}
	return ranker
}

func main() {
	data, err := ReadFileToData("./data/ml-100k/u.data")
	if err != nil {
		fmt.Println(err)
	}

	encountered, userItemMatrix, err := preprocessing.MakeUserItemMatrix(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println("+++User Simlarity+++")

	fmt.Println("")
	fmt.Println("Pearson")
	fmt.Println("-----------------------------")
	simMat := MakeSimilarityMatrix(userItemMatrix, Pearson, true)
	MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	// fmt.Println("")
	// fmt.Println("Cosine")
	// fmt.Println("-----------------------------")
	// simMat = MakeSimilarityMatrix(userItemMatrix, Cosine, true)
	// MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	// MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	// fmt.Println("")
	// fmt.Println("Adjusted Cosine")
	// fmt.Println("-----------------------------")
	// simMat = MakeSimilarityMatrix(userItemMatrix, AdjustedCosine, true)
	// MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	// MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	fmt.Println("")
	fmt.Println("+++Item Simlarity+++")

	// fmt.Println("")
	// fmt.Println("Pearson")
	// fmt.Println("-----------------------------")
	// simMat = MakeSimilarityMatrix(userItemMatrix, Pearson, false)
	// MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	// MostSimilar(encountered.UniqueItem, simMat, 1501, 3)

	// fmt.Println("")
	// fmt.Println("Cosine")
	// fmt.Println("-----------------------------")
	// simMat = MakeSimilarityMatrix(userItemMatrix, Cosine, false)
	// MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	// MostSimilar(encountered.UniqueItem, simMat, 1501, 3)

	fmt.Println("")
	fmt.Println("Adjusted Cosine")
	fmt.Println("-----------------------------")
	simMat = MakeSimilarityMatrix(userItemMatrix, AdjustedCosine, false)
	MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	MostSimilar(encountered.UniqueItem, simMat, 1501, 3)

}
