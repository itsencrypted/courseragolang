package main

import (
"fmt"
"encoding/json"
)

type Test struct {
	Name string
	Address string
}

func main(){
	entry:= Test{
		Name: "Citibank",
		Address: "Av. Paulista 1111, Sao Paulo, SP, Brazil"}

		ret, err := json.Marshal(entry)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println(string(ret))
	}



