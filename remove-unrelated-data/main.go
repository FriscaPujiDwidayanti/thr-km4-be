package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
    return dataMap

}

func main() {
	datMap := map[string]interface{}{
	 	"name":    "Edo",
	 	"age":     20,
	 	"address": "Jakarta",
	}
	fmt.Println(datMap)

	keyDelete := "address"
	newdatMap := removeUnrelated(datMap, keyDelete)
	fmt.Println(newdatMap)
}
