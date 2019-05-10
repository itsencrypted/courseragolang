package  main

import ("fmt"
"bufio"
"os"
"strings"
"strconv"
"sort")


func printSlice(s []int){
	fmt.Printf("%v\n", s)
}

func main(){
	var intSlice = make([]int,0,3)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("please enter an integer: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	fmt.Printf("user sent in %s\n",userInput)

	for userInput != "X" {
		intVal,err := strconv.ParseInt(userInput,0,0)
		if err != nil {
			fmt.Printf("err encountered %s\n",err.Error())
		}
		intSlice = append(intSlice,int(intVal))
		sort.Ints(intSlice)
		fmt.Print("Please enter an integer: ")
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
	}
}

