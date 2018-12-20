package main_test

import (
	"log"
	"testing"

	"github.com/lapis-zero09/GoRec/pkg/matrix"
	"github.com/lapis-zero09/GoRec/pkg/preprocess"
	"github.com/lapis-zero09/GoRec/pkg/similarity"
)

const file = "../../data/ml-100k/u.data"

var sb matrix.SimBox

func init() {
	dPtr, err := preprocess.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	sb.Data = dPtr
	if err := sb.MakeUserItemMatrix(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkPearsonUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Pearson, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkCosineUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Cosine, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkAdjustedCosineUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.AdjustedCosine, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkJaccardUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Jaccard, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkDiceUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Dice, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkSimpsonUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Simpson, true); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	}
}

func BenchmarkPearsonItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Pearson, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}

func BenchmarkCosineItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Cosine, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}

func BenchmarkAdjustedCosineItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.AdjustedCosine, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}

func BenchmarkJaccardItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Jaccard, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}

func BenchmarkDiceItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Dice, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}

func BenchmarkSimpsonItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := sb.MakeSimilarityMatrix(similarity.Simpson, false); err != nil {
			log.Fatal(err)
		}
		_ = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	}
}
