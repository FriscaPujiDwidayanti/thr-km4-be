package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
        return arr[1:]
    } else if position == "right" {
        return arr[:len(arr)-1]
    } else {
        return arr
    }

	//return nil
}

func main() {
    arr1 := []interface{}{1, 2, 3, 4, 5}
    newArr1 := removeLeftRight(arr1, "left")
    fmt.Println(newArr1) 

    arr2 := []interface{}{1, 2, 3, 4, 5}
    newArr2 := removeLeftRight(arr2, "right")
    fmt.Println(newArr2) 

    arr3 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    newArr3 := removeLeftRight(arr3, "left")
    fmt.Println(newArr3) 

    arr4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
    newArr4 := removeLeftRight(arr4, "right")
    fmt.Println(newArr4) 
}
