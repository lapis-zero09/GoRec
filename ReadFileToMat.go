package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	// "reflect"
)

func StringToInt(string_slice []string) []int {
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

func ReadFileToMat(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file can't be opened")
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'

	data_mat := [][]int{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("file can't be readed")
			return nil, fmt.Errorf("file can't be opened")
		}
		data := StringToInt(record[:3])
		data_mat = append(data_mat, data)
	}

	return data_mat, nil
}

func main() {
	data, err := ReadFileToMat("./data/ml-100k/u.data")
	if err != nil {
		fmt.Errorf("something went wrong in readFile")
	}
	fmt.Println(data[0])
}
