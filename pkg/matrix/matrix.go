package matrix

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sort"

	"github.com/lapis-zero09/GoRec/pkg/cast"
)

type DataFrame [][]int

type Encountered struct {
	UniqueUser []int
	UniqueItem []int
}

func CountTrueSize(bool_arr []bool) int {
	var trueSize int
	for _, b := range bool_arr {
		if b {
			trueSize++
		}
	}
	return trueSize
}

func (data_df DataFrame) shape() (int, int) {
	return len(data_df), len(data_df[0])
}

func find(arr []int, val int) (int, error) {
	for i, unique_id := range arr {
		if unique_id == val {
			return i, nil
		}
	}
	return 0, fmt.Errorf("can't find that id in input arr")
}

func (data_df DataFrame) ReturnUniqueSize(encountered Encountered) ([]int, Encountered) {
	user := make([]bool, len(data_df))
	item := make([]bool, len(data_df))

	for _, val := range data_df {
		// user
		if !user[val[0]] {
			user[val[0]] = true
			encountered.UniqueUser = append(encountered.UniqueUser, val[0])
		}

		// item
		if !item[val[1]] {
			item[val[1]] = true
			encountered.UniqueItem = append(encountered.UniqueItem, val[1])
		}
	}
	sort.Ints(encountered.UniqueUser)
	sort.Ints(encountered.UniqueItem)
	uniqueSize := []int{CountTrueSize(user), CountTrueSize(item)}
	return uniqueSize, encountered
}

func MakeUserItemMatrix(d [][]int) (Encountered, DataFrame, error) {
	var data DataFrame = DataFrame(d)
	encountered := Encountered{}
	uniqueSize, encountered := data.ReturnUniqueSize(encountered)
	UniqueUserSize, UniqueItemSize := uniqueSize[0], uniqueSize[1]

	df := make(DataFrame, UniqueUserSize)
	for i := range df {
		df[i] = make([]int, UniqueItemSize)
	}

	for _, val := range data {
		user_id, err := find(encountered.UniqueUser, val[0])
		if err != nil {
			fmt.Println(err)
		}
		item_id, err := find(encountered.UniqueItem, val[1])
		if err != nil {
			fmt.Println(err)
		}
		df[user_id][item_id] = val[2]
	}
	fmt.Println("UserSize\tItemSize")
	fmt.Println("-----------------------------")
	l, m := df.shape()
	fmt.Printf(" %d \t %d \n", l, m)
	fmt.Println("-----------------------------")
	return encountered, df, nil
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
	log.Println(funcName)
	log.Println(method)
	if funcName == "github.com/lapis-zero09/GoRec/pkg/matrix.AdjustedCosine" {
		if userFlag {
			mean = make([]float64, len(userItemMatrix[0]))
			for i := 0; i < len(userItemMatrix[0]); i++ {
				tmpVec, err := TakeCol(userItemMatrix, i)
				if err != nil {
					fmt.Println(err)
					break
				}
				mean[i] = Mean(cast.IntSliceToFloatSlice(tmpVec))
			}
		} else {
			mean = make([]float64, len(userItemMatrix))
			for i, val := range userItemMatrix {
				mean[i] = Mean(cast.IntSliceToFloatSlice(val))
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

func MostSimilar(encounteredUnique []int, simMat [][]float64, id, similarSize int) {
	simVector := simMat[id]
	sortedSimVector := make([]float64, len(simVector))
	copy(sortedSimVector, simVector)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedSimVector)))

	fmt.Printf("mainId = %d\n", id)
	fmt.Println("rank\tid\t similarity")
	fmt.Println("-----------------------------")
	rank := 1
	for _, sortedSim := range sortedSimVector[:similarSize] {
		for i, sim := range simVector {
			if sim == sortedSim {
				fmt.Printf(" %d  \t %d \t %f\n", rank, encounteredUnique[i], sim)
				rank++
			}
		}
	}
	fmt.Println("-----------------------------")
}
