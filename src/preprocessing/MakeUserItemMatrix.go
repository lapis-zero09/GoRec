package preprocessing

import (
	"fmt"
	"sort"
)

type DataFrame [][]int

type Encountered struct {
	uniqueUser []int
	uniqueItem []int
}

func (data_df DataFrame) TakeCol(colidx int) ([][]int, error) {
	t := [][]int{}
	if colidx < 0 || len(data_df) < colidx {
		return nil, fmt.Errorf("column index is invalid value!")
	}
	return t, nil
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
			encountered.uniqueUser = append(encountered.uniqueUser, val[0])
		}

		// item
		if !item[val[1]] {
			item[val[1]] = true
			encountered.uniqueItem = append(encountered.uniqueItem, val[1])
		}
	}
	sort.Ints(encountered.uniqueUser)
	sort.Ints(encountered.uniqueItem)
	uniqueSize := []int{CountTrueSize(user), CountTrueSize(item)}
	return uniqueSize, encountered
}

func MakeUserItemMatrix(d [][]int) (DataFrame, error) {
	var data DataFrame = DataFrame(d)

	// user_id, err := data.TakeCol(0)
	// if err != nil {
	// 	fmt.Errorf("something went wrong in TakeCol")
	// }
	// item_id, err := data.TakeCol(1)
	// if err != nil {
	// 	fmt.Errorf("something went wrong in TakeCol")
	// }
	// fmt.Println(user_id[0], item_id[0])

	encountered := Encountered{}
	uniqueSize, encountered := data.ReturnUniqueSize(encountered)
	uniqueUserSize, uniqueItemSize := uniqueSize[0], uniqueSize[1]

	df := make(DataFrame, uniqueUserSize)
	for i := range df {
		df[i] = make([]int, uniqueItemSize)
	}

	// f := make(DataFrame, 5)
	// for i := range f {
	// 	f[i] = make([]int, 3)
	// }
	// fmt.Println(f.shape())
	// fmt.Println(df[0])

	for _, val := range data {
		user_id, err := find(encountered.uniqueUser, val[0])
		if err != nil {
			fmt.Println(err)
		}
		item_id, err := find(encountered.uniqueItem, val[1])
		if err != nil {
			fmt.Println(err)
		}
		df[user_id][item_id] = val[2]
	}
	fmt.Println(df.shape())
	return df, nil
}
