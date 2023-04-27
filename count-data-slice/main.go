package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
}

func main() {
	elementData1 := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	elementData2 := []any{"Edo", "Budi", "Joko", "Tono"}
    fmt.Println(howManyElements(elementData1))
	fmt.Println(howManyElements(elementData2))
	
}
