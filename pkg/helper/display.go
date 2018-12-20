package helper

import "fmt"

func Print(ranker map[int]float64) {
	fmt.Println("rank\tid\t similarity")
	fmt.Println("-----------------------------")
	rank := 1
	for k, v := range ranker {
		fmt.Printf(" %d  \t %d \t %f\n", rank, k, v)
		rank++
	}
	fmt.Println("-----------------------------")
}
