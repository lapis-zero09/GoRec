package main

import (
	"encoding/csv"
	"fmt"
	"gorec/src/preprocessing"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
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

func StrsliceToFloatslice(string_slice []int) []float64 {
	int_array := make([]int, cap(string_slice))
	for i, val := range string_slice {
		val, err := float64(val)
		if err != nil {
			fmt.Println(err)
		}
		int_array[i] = val
	}
	return int_array
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
	// lines := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't found EOF")
		}
		// lines += 1
		data_int := StrsliceToIntslice(record[:3])
		data = append(data, data_int)
	}
	// fmt.Printf("lines: %d", lines)
	return data, nil
}

func Mean(arr []int) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += float64(val)
	}
	return sum / float64(len(arr))
}

func SubMean(arr []int) []float64 {
	mean := Mean(arr)
	diff_arr := make([]float64, len(arr))
	for i, val := range arr {
		diff_arr[i] = float64(val) - mean
	}
	return diff_arr
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

func SumSquad(arr []float64) float64 {
	var res float64
	for _, val := range arr {
		res += (val * val)
	}
	return res
}

func Cosine(userVec1, userVec2 []int) (float64, error) {
	numerator, err := Dot(userVec1, userVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(userVec1)) * math.Sqrt(SumSquad(userVec2))
	return (numerator / deliminator), nil
}

func Pearson(userVec1, userVec2 []int) (float64, error) {
	u_diff := SubMean(userVec1)
	v_diff := SubMean(userVec2)
	numerator, err := Dot(u_diff, v_diff)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(u_diff)) * math.Sqrt(SumSquad(v_diff))
	return (numerator / deliminator), nil
}

func MostSimilarUser(encountered preprocessing.Encountered, userItemMatrix [][]int, userId, similarSize int, method string) {
	userSize := len(userItemMatrix)
	userSimMat := make([][]float64, userSize)
	for i := 0; i < userSize; i++ {
		userSimMat[i] = make([]float64, userSize)
	}

	for i := 0; i < userSize; i++ {
		for j := userSize - 1; j > i; j-- {
			if method == "pearson" {
				sim, err := Pearson(userItemMatrix[i], userItemMatrix[j])
				if err != nil {
					fmt.Println(err)
					break
				}
			} else if method == "cosine" {
				sim, err := Cosine(userItemMatrix[i], userItemMatrix[j])
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			userSimMat[i][j] = sim
			userSimMat[j][i] = sim
		}
	}

	userSimVector := userSimMat[userId]
	sortedUserSimVector := make([]float64, len(userSimVector))
	copy(sortedUserSimVector, userSimVector)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedUserSimVector)))

	fmt.Println("rank userId similarity")
	rank := 1
	for i, sim := range userSimVector {
		for _, sortedSim := range sortedUserSimVector[:similarSize] {
			if sim == sortedSim {
				fmt.Println(rank, encountered.UniqueUser[i], sim)
				rank++
			}
		}
	}
}

// func TakeCol(data [][]int, colidx int) ([][]int, error) {
// 	t := [][]int{}
// 	if colidx < 0 || len(data) < colidx {
// 		return nil, fmt.Errorf("column index is invalid value!")
// 	}
// 	return t, nil
// }

func main() {
	data, err := ReadFileToData("./data/ml-100k/u.data")
	if err != nil {
		fmt.Println(err)
	}

	encountered, userItemMatrix, err := preprocessing.MakeUserItemMatrix(data)
	if err != nil {
		fmt.Println(err)
	}

	MostSimilarUser(encountered, userItemMatrix, 941, 3, "pearson")
	MostSimilarUser(encountered, userItemMatrix, 941, 3, "cosine")
}
