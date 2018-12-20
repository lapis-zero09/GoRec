package preprocess

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/lapis-zero09/GoRec/pkg/matrix"

	"github.com/lapis-zero09/GoRec/pkg/cast"
)

func ReadFile(filename string) (*matrix.DataFrame, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file can't be opened from os open")
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'
	var data matrix.DataFrame
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can't found EOF")
		}
		d := cast.StrSliceToIntslice(record[:3])
		data = append(data, d)
	}
	return &data, nil
}
