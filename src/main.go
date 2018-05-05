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

func Mean(arr []float64) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += float64(val)
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

func Cosine(userVec1, userVec2 []int) (float64, error) {
	tmpUserVec1 := IntsliceToFloatslice(userVec1)
	tmpUserVec2 := IntsliceToFloatslice(userVec2)
	numerator, err := Dot(tmpUserVec1, tmpUserVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(tmpUserVec1)) * math.Sqrt(SumSquad(tmpUserVec2))
	return (numerator / deliminator), nil
}

func Jaccard(userVec1, userVec2 []int) (float64, error) {
	tmpUserVec1 := IntsliceToFloatslice(userVec1)
	tmpUserVec2 := IntsliceToFloatslice(userVec2)
	numerator, err := Dot(tmpUserVec1, tmpUserVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := (Sum(tmpUserVec1) + Sum(tmpUserVec2)) - numerator
	return (numerator / deliminator), nil
}

func Dice(userVec1, userVec2 []int) (float64, error) {
	tmpUserVec1 := IntsliceToFloatslice(userVec1)
	tmpUserVec2 := IntsliceToFloatslice(userVec2)
	numerator, err := Dot(tmpUserVec1, tmpUserVec2)
	if err != nil {
		return 0.0, err
	}
	deliminator := Sum(tmpUserVec1) + Sum(tmpUserVec2)
	return (2 * numerator / deliminator), nil
}

func Simpson(userVec1, userVec2 []int) (float64, error) {
	tmpUserVec1 := IntsliceToFloatslice(userVec1)
	tmpUserVec2 := IntsliceToFloatslice(userVec2)
	numerator, err := Dot(tmpUserVec1, tmpUserVec2)
	if err != nil {
		return 0.0, err
	}
	sum := Sum(tmpUserVec1)
	if tmpSum := Sum(tmpUserVec2); sum > tmpSum {
		return (numerator / tmpSum), nil
	}
	return (numerator / sum), nil
}

func Pearson(userVec1, userVec2 []int) (float64, error) {
	tmpUserVec1 := IntsliceToFloatslice(userVec1)
	tmpUserVec2 := IntsliceToFloatslice(userVec2)
	u_diff := SubMean(tmpUserVec1)
	v_diff := SubMean(tmpUserVec2)
	numerator, err := Dot(u_diff, v_diff)
	if err != nil {
		return 0.0, err
	}
	deliminator := math.Sqrt(SumSquad(u_diff)) * math.Sqrt(SumSquad(v_diff))
	return (numerator / deliminator), nil
}

func MakeSimilarityMatrix(userItemMatrix [][]int, method func([]int, []int) (float64, error)) [][]float64 {
	userSize := len(userItemMatrix)
	userSimMat := make([][]float64, userSize)
	for i := 0; i < userSize; i++ {
		userSimMat[i] = make([]float64, userSize)
	}

	for i := 0; i < userSize; i++ {
		for j := userSize - 1; j > i; j-- {
			sim, err := method(userItemMatrix[i], userItemMatrix[j])
			if err != nil {
				fmt.Println(err)
				break
			}
			userSimMat[i][j] = sim
			userSimMat[j][i] = sim
		}
	}
	return userSimMat
}

func MostSimilarUser(encountered preprocessing.Encountered, userSimMat [][]float64, userId, similarSize int) {
	userSimVector := userSimMat[userId]
	sortedUserSimVector := make([]float64, len(userSimVector))
	copy(sortedUserSimVector, userSimVector)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedUserSimVector)))

	fmt.Printf("mainUserId = %d\n", userId)
	fmt.Println("rank\tuserId\tsimilarity")
	fmt.Println("-----------------------------")
	rank := 1
	for _, sortedSim := range sortedUserSimVector[:similarSize] {
		for i, sim := range userSimVector {
			if sim == sortedSim {
				fmt.Printf(" %d  \t %d  \t %f\n", rank, encountered.UniqueUser[i], sim)
				rank++
			}
		}
	}
	fmt.Println("-----------------------------")
}

// func TakeCol(data [][]int, colidx int) ([][]int, error) {
//     t := [][]int{}
//     if colidx < 0 || len(data) < colidx {
//         return nil, fmt.Errorf("column index is invalid value!")
//     }
//     return t, nil
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

	// fmt.Println(encountered, userItemMatrix)
	fmt.Println("")
	fmt.Println("Pearson")
	fmt.Println("-----------------------------")
	userSimMat := MakeSimilarityMatrix(userItemMatrix, Pearson)
	MostSimilarUser(encountered, userSimMat, 941, 2)
	MostSimilarUser(encountered, userSimMat, 356, 3)

	fmt.Println("")
	fmt.Println("Cosine")
	fmt.Println("-----------------------------")
	userSimMat = MakeSimilarityMatrix(userItemMatrix, Cosine)
	MostSimilarUser(encountered, userSimMat, 941, 2)
	MostSimilarUser(encountered, userSimMat, 356, 3)
}
