package main

import (
	"fmt"
	"strings"
	"strconv"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	var userData User
    fields := strings.Split(data, ",")
    userData.Name = fields[0]
    userData.Age, _ = strconv.Atoi(fields[1])
    userData.Address = fields[2]
    return userData
	
	//return User{}
}

func main() {
	data := "Edo,20,Jakarta"
	userData := ConvertData(data)
	fmt.Printf("%+v\n", userData)
	//fmt.Println("{ name:", userData.Name + "address:", userData.Address+"}")
	//fmt.Println({"name:" +userData.Name+, age: +userData.Age, "address: ", userData.Address+})
}
