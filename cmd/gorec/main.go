package main

import (
	"fmt"
	"log"

	"github.com/lapis-zero09/GoRec/pkg/matrix"
	"github.com/lapis-zero09/GoRec/pkg/preprocess"
)

type Encountered struct {
	UniqueUser []int
	UniqueItem []int
}

func main() {
	d, err := preprocess.ReadFile("./data/ml-100k/u.data")
	if err != nil {
		log.Fatal(err)
	}

	encountered, userItemMatrix, err := matrix.MakeUserItemMatrix(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("+++User Simlarity+++")

	fmt.Println("")
	fmt.Println("Pearson")
	fmt.Println("-----------------------------")
	simMat := matrix.MakeSimilarityMatrix(userItemMatrix, matrix.Pearson, true)
	matrix.MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	matrix.MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	// fmt.Println("")
	// fmt.Println("Cosine")
	// fmt.Println("-----------------------------")
	// simMat = matrix.MakeSimilarityMatrix(userItemMatrix, Cosine, true)
	// matrix.MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	// matrix.MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	// fmt.Println("")
	// fmt.Println("Adjusted Cosine")
	// fmt.Println("-----------------------------")
	// simMat = matrix.MakeSimilarityMatrix(userItemMatrix, AdjustedCosine, true)
	// matrix.MostSimilar(encountered.UniqueUser, simMat, 941, 2)
	// matrix.MostSimilar(encountered.UniqueUser, simMat, 356, 3)

	fmt.Println("")
	fmt.Println("+++Item Simlarity+++")

	// fmt.Println("")
	// fmt.Println("Pearson")
	// fmt.Println("-----------------------------")
	// simMat = matrix.MakeSimilarityMatrix(userItemMatrix, Pearson, false)
	// matrix.MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	// matrix.MostSimilar(encountered.UniqueItem, simMat, 1501, 3)

	// fmt.Println("")
	// fmt.Println("Cosine")
	// fmt.Println("-----------------------------")
	// simMat = matrix.MakeSimilarityMatrix(userItemMatrix, Cosine, false)
	// matrix.MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	// matrix.MostSimilar(encountered.UniqueItem, simMat, 1501, 3)

	fmt.Println("")
	fmt.Println("Adjusted Cosine")
	fmt.Println("-----------------------------")
	simMat = matrix.MakeSimilarityMatrix(userItemMatrix, matrix.AdjustedCosine, false)
	matrix.MostSimilar(encountered.UniqueItem, simMat, 941, 2)
	matrix.MostSimilar(encountered.UniqueItem, simMat, 1501, 3)
}
