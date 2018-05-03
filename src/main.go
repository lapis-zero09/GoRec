package main

import (
	"encoding/csv"
	"fmt"
	"gorec/src/preprocessing"
	"io"
	"os"
	"strconv"
)

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

func main() {
	data, err := ReadFileToData("./data/ml-100k/u.data")
	if err != nil {
		fmt.Println(err)
	}

	userItemMatrix, err := preprocessing.MakeUserItemMatrix(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userItemMatrix[25])
	}

	// uesr_size := len(userItemMatrix)

}
