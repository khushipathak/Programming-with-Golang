package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var sli []int
	fmt.Println("Enter a number to append to the slice, enter a character to exit the program: ")

	for scanner.Scan() {
		n := scanner.Text()

		if n, err := strconv.Atoi(n); err == nil {
		
			sli = append(sli, n)
		
			sort.SliceStable(sli, func(i, j int) bool { return sli[i] < sli[j] })
			fmt.Println("Sorted slice:", sli)
		} else {
			break;
		}

	}
	

}
