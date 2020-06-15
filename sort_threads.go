package main

import (
	"fmt"
	"sync"
	"sort"
)

func sort_slice1(sli []int, wg *sync.WaitGroup){

	sort.Ints(sli)
	defer wg.Done()

}

func sort_slice2(sli []int, wg *sync.WaitGroup){

	sort.Ints(sli)
	defer wg.Done()

}

func sort_slice3(sli []int, wg *sync.WaitGroup){

	sort.Ints(sli)
	defer wg.Done()

}

func sort_slice4(sli []int, wg *sync.WaitGroup){

	sort.Ints(sli)
	defer wg.Done()

}

func merge(sli1 []int, sli2 []int) []int {

	res := append(sli1, sli2...)
	sort.Ints(res)
	return res

}

func main() {

	var n,x,cut int
	var wg sync.WaitGroup

	fmt.Println("--------------------------------")

	fmt.Println("Enter the number of integers to sort:")
	fmt.Scan(&n)
	sli := make([]int, 0, n)

	fmt.Println("Enter the integers to sort:")
	for i := 0; i < n; i++{
		fmt.Scan(&x)
		sli = append(sli, x)
	}
	fmt.Println("The slice you entered is: ", sli)

	cut = n/4

	sli1 := sli[:cut]
	sli2 := sli[cut:2*cut]
	sli3 := sli[2*cut:3*cut]
	sli4 := sli[3*cut:]

	wg.Add(4)
	sort_slice1(sli1, &wg)
	sort_slice2(sli2, &wg)
	sort_slice3(sli3, &wg)
	sort_slice4(sli4, &wg)
	wg.Wait()

	sli_merge_1 := merge(sli1,sli2)
	sli_merge_2 := merge(sli3,sli4)

	sli_final := merge(sli_merge_1,sli_merge_2)

	fmt.Println("The sorted slice is: ", sli_final)

}