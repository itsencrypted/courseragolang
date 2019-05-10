package main

import "fmt"

func main() {
	var sli []int
	printSlice(sli)

	sli = append (sli,0)
	printSlice(sli)

	sli = append (sli, 3)
	printSlice(sli)

	sli = append (sli, 2, 3, 4)
	printSlice(sli)
}

func printSlice(sli []int){
	fmt.Printf("len=%d cap=%d %v\n", len(sli), cap(sli), sli)
}





