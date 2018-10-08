package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func swap(arr []string, first int, second int) {
	tmp := arr[first]
	arr[first] = arr[second]
	arr[second] = tmp
}

func separator(arr []string, left int, right int) int {
	var leader int
	leader, _ = strconv.Atoi(arr[(right+left)/2])
	for left <= right {
		var i int
		for i, _ = strconv.Atoi(arr[left]); i < leader; {
			left++
			i, _ = strconv.Atoi(arr[left])
		}

		var j int
		for j, _ = strconv.Atoi(arr[right]); j > leader; {
			right--
			j, _ = strconv.Atoi(arr[right])
		}

		if left <= right {
			swap(arr, left, right)
			left++
			right--
		}
	}
	return left
}

func quicksort(arr []string, left int, right int) {
	if left < right {
		p := separator(arr, left, right)
		quicksort(arr, left, p-1)
		quicksort(arr, p, right)
	}
}

func main() {
	// Самым первым аргументом всегда идет имя бинарника
	if len(os.Args) == 1 {
		fmt.Println("File doesn't specified")
		fmt.Println("Use:")
		fmt.Println(os.Args[0] + " [filename]")
		os.Exit(0)
	}

	fname := os.Args[1]
	myBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error occured while reading file")
	}
	str := string(myBytes)
	numbers := strings.Split(str, " ")
	quicksort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}
