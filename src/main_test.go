package main

import (
	"log"
	"testing"

	"github.com/lapis-zero09/GoRec/src/preprocessing"
)

const file = "./data/ml-100k/u.data"

var En preprocessing.Encountered
var UserItemMatrix [][]int

func init() {
	data, err := ReadFileToData(file)
	if err != nil {
		log.Fatal(err)
	}

	En, UserItemMatrix, err = preprocessing.MakeUserItemMatrix(data)
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkPearsonUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Pearson, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}

func BenchmarkCosineUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Cosine, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}

func BenchmarkAdjustedCosineUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, AdjustedCosine, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}

func BenchmarkJaccardUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Jaccard, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}

func BenchmarkDiceUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Dice, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}

func BenchmarkSimpsonUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Simpson, true)
		_ = MostSimilar(En.UniqueUser, simMat, 941, 2)
	}
}
func BenchmarkPearsonItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Pearson, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}

func BenchmarkCosineItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Cosine, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}

func BenchmarkAdjustedCosineItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, AdjustedCosine, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}

func BenchmarkJaccardItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Jaccard, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}

func BenchmarkDiceItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Dice, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}

func BenchmarkSimpsonItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simMat := MakeSimilarityMatrix(UserItemMatrix, Simpson, false)
		_ = MostSimilar(En.UniqueItem, simMat, 941, 2)
	}
}
