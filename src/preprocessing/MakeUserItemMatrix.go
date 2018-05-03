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
	encountered := Encountered{}
	uniqueSize, encountered := data.ReturnUniqueSize(encountered)
	uniqueUserSize, uniqueItemSize := uniqueSize[0], uniqueSize[1]

	df := make(DataFrame, uniqueUserSize)
	for i := range df {
		df[i] = make([]int, uniqueItemSize)
	}

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
	fmt.Println("UserSize ItemSize")
	fmt.Println(df.shape())
	return df, nil
}
