package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lapis-zero09/GoRec/pkg/helper"
	"github.com/lapis-zero09/GoRec/pkg/matrix"
	"github.com/lapis-zero09/GoRec/pkg/preprocess"
	"github.com/lapis-zero09/GoRec/pkg/similarity"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify data file name as argument.")
	}
	var sb matrix.SimBox
	dPtr, err := preprocess.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sb.Data = dPtr

	if err := sb.MakeUserItemMatrix(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("+++User Simlarity+++")

	fmt.Println("")
	fmt.Println("Pearson")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Pearson, true); err != nil {
		log.Fatal(err)
	}
	ranker := sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Cosine")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Cosine, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Adjusted Cosine")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.AdjustedCosine, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Jaccard")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Jaccard, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Dice")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Dice, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Simpson")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Simpson, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("+++Item Simlarity+++")

	fmt.Println("")
	fmt.Println("Pearson")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Pearson, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Cosine")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Cosine, true); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueUser, 356, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Adjusted Cosine")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.AdjustedCosine, false); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 1501, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Jaccard")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Jaccard, false); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 1501, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Dice")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Dice, false); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 1501, 3)
	helper.Print(ranker)

	fmt.Println("")
	fmt.Println("Simpson")
	fmt.Println("-----------------------------")
	if err := sb.MakeSimilarityMatrix(similarity.Simpson, false); err != nil {
		log.Fatal(err)
	}
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 941, 2)
	helper.Print(ranker)
	ranker = sb.MostSimilar(sb.Uniques.UniqueItem, 1501, 3)
	helper.Print(ranker)
}
