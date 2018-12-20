package matrix

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sort"

	"github.com/lapis-zero09/GoRec/pkg/cast"
	"github.com/lapis-zero09/GoRec/pkg/similarity"
)

type DataFrame [][]int
type FloatDataFrame [][]float64

type Uniques struct {
	UniqueUser []int
	UniqueItem []int
}

type SimBox struct {
	Data, Matrix     *DataFrame
	SimilarityMatrix *FloatDataFrame
	Uniques          *Uniques
}

func uniqueSize(bool_arr []bool) int {
	var trueSize int
	for _, b := range bool_arr {
		if b {
			trueSize++
		}
	}
	return trueSize
}

func (d DataFrame) shape() (int, int) {
	return len(d), len(d[0])
}

func find(arr []int, val int) (int, error) {
	for i, unique_id := range arr {
		if unique_id == val {
			return i, nil
		}
	}
	return 0, fmt.Errorf("can't find that id in input arr")
}

func (d *DataFrame) TakeCol(colIdx int) ([]int, error) {
	if colIdx < 0 || len((*d)[0]) < colIdx {
		return nil, fmt.Errorf("column index is invalid value!")
	}
	a := make([]int, len((*d)))
	for i, val := range *d {
		a[i] = val[colIdx]
	}
	return a, nil
}

func (sb *SimBox) Unique() (int, int) {
	u := &Uniques{}
	user := make([]bool, len(*sb.Data))
	item := make([]bool, len(*sb.Data))

	for _, val := range *sb.Data {
		// user
		if !user[val[0]] {
			user[val[0]] = true
			u.UniqueUser = append(u.UniqueUser, val[0])
		}

		// item
		if !item[val[1]] {
			item[val[1]] = true
			u.UniqueItem = append(u.UniqueItem, val[1])
		}
	}
	sort.Ints(u.UniqueUser)
	sort.Ints(u.UniqueItem)
	sb.Uniques = u
	return uniqueSize(user), uniqueSize(item)
}

func (sb *SimBox) MakeUserItemMatrix() error {
	userSize, itemSize := sb.Unique()
	log.Printf("UserSize: %d", userSize)
	log.Printf("ItemSize: %d", itemSize)

	df := make(DataFrame, userSize)
	for i := range df {
		df[i] = make([]int, itemSize)
	}

	for _, val := range *sb.Data {
		userID, err := find(sb.Uniques.UniqueUser, val[0])
		if err != nil {
			return err
		}
		itemID, err := find(sb.Uniques.UniqueItem, val[1])
		if err != nil {
			return err
		}
		df[userID][itemID] = val[2]
	}

	sb.Matrix = &df
	return nil
}

func (sb *SimBox) MakeSimilarityMatrix(method func([]int, []int, ...[]float64) (float64, error), userFlag bool) error {
	var size int
	if userFlag {
		size, _ = sb.Matrix.shape()
	} else {
		_, size = sb.Matrix.shape()
	}

	simMat := make(FloatDataFrame, size)
	for i := 0; i < size; i++ {
		simMat[i] = make([]float64, size)
	}

	var mean []float64
	if funcName := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name(); funcName == "github.com/lapis-zero09/GoRec/pkg/similarity.AdjustedCosine" {
		if userFlag {
			mean = make([]float64, len((*sb.Matrix)[0]))
			for i := 0; i < len((*sb.Matrix)[0]); i++ {
				tmpVec, err := sb.Matrix.TakeCol(i)
				if err != nil {
					return err
				}
				mean[i] = similarity.Mean(cast.IntSliceToFloatSlice(tmpVec))
			}
		} else {
			mean = make([]float64, len((*sb.Matrix)))
			for i, val := range *sb.Matrix {
				mean[i] = similarity.Mean(cast.IntSliceToFloatSlice(val))
			}
		}
	}
	var err error
	var sim float64
	for i := 0; i < size; i++ {
		for j := size - 1; j > i; j-- {
			if userFlag {
				sim, err = method((*sb.Matrix)[i], (*sb.Matrix)[j], mean)
			} else {
				iVec, err := sb.Matrix.TakeCol(i)
				jVec, err := sb.Matrix.TakeCol(j)
				if err != nil {
					return err
				}
				sim, err = method(iVec, jVec, mean)
			}
			if err != nil {
				return err
			}
			simMat[i][j] = sim
			simMat[j][i] = sim
		}
	}

	sb.SimilarityMatrix = &simMat
	return nil
}

func (sb *SimBox) MostSimilar(u []int, id, similarSize int) map[int]float64 {
	simVector := (*sb.SimilarityMatrix)[id]
	sortedSimVector := make([]float64, len(simVector))
	copy(sortedSimVector, simVector)
	sort.Sort(sort.Reverse(sort.Float64Slice(sortedSimVector)))

	log.Printf("target ID: %d", id)
	ranker := make(map[int]float64, similarSize)
	for _, sortedSim := range sortedSimVector[:similarSize] {
		for i, sim := range simVector {
			if sim == sortedSim {
				ranker[u[i]] = sim
			}
		}
	}
	return ranker
}
