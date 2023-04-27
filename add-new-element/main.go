package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...) 
	} else if position == "down" {
		return append(data, newData)
	} else {
		return data
	}
	//return nil
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	newDat1 := 6
	position1 := "up"
	newArr1 := AddElement(arr1, newDat1, position1)
	fmt.Println(newArr1)
	
	arr2 := []int{1, 2, 3, 4, 5}
	newDat2 := 6
	position2 := "down"
	newArr2 := AddElement(arr2, newDat2, position2)
	fmt.Println(newArr2)
}
