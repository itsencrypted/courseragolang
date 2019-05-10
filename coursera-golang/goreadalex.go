package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	var filename string
	var filedata *[]byte
	structs := make([]Name, 0)

	// while loop until we successfully read file
	for {
		//get filename from user
		fmt.Printf("Enter filename: ")
		fmt.Scan(&filename)

		//try reading file
		data, err := ioutil.ReadFile(filename)

		if err == nil {
			filedata = &data
			break
		}
		
		fmt.Println("Error reading file, make sure filename is correct!")
	}

	//replace newlines to ensure it works both on win and unix
	names := strings.Split(strings.Replace(string(*filedata), "\r\n", "\n", -1), "\n")

	//save names as structs
	for _, v := range names {
		if v == "" {
			//there can be trailing empty 'name' because of final new line in file continue
		}
		nameParts := strings.Split(v, "")
		structs = append(structs, Name{Fname: nameParts[0], Lname: nameParts[1]})
	}

	//print all names
	fmt.Printf("\nList of names found in file %v:\n\n", filename)
	for _, v := range structs {
		fmt.Printf("%v %v\n", v.Fname, v.Lname)
	}
}

