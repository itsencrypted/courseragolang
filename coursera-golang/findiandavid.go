package main
import (
	"fmt"
	"strings"
)

func main(){
	var s string
	fmt.Println("Please enter a string")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Input was not a string. Please rerun the program and input a string next time")
	}

	s = strings.ToLower(s)
	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n"){
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

