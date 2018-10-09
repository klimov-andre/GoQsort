package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func separator(arr []int, left int, right int) int {
	var leader int
	leader = arr[(right+left)/2]
	for left <= right {
		var i int
		for i = arr[left]; i < leader; {
			left++
			i = arr[left]
		}
		var j int
		for j = arr[right]; j > leader; {
			right--
			j = arr[right]
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func quicksort(arr []int, left int, right int) {
	if left < right {
		p := separator(arr, left, right)
		quicksort(arr, left, p-1)
		quicksort(arr, p, right)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File doesn't specified")
		fmt.Println("Use:")
		fmt.Println(os.Args[0] + " [filename]")
		os.Exit(0)
	}

	fname := os.Args[1]
	myBytes, err := ioutil.ReadFile(fname)
	var numbers []int
	var sortedNumbers []int

	if err != nil {
		fmt.Println("Error occured while reading file:")
		fmt.Println(err)
	} else {
		str := string(myBytes)
		numbersString := strings.Split(str, " ")

		for i := 0; i < len(numbersString); i++ {
			num, err := strconv.Atoi(numbersString[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			numbers = append(numbers, num)
		}
		sortedNumbers = make([]int, len(numbers))
		copy(sortedNumbers, numbers)
		quicksort(sortedNumbers, 0, len(sortedNumbers)-1)
		fmt.Println(sortedNumbers)
	}
}
