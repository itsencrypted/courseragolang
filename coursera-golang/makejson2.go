package main

import (
	"fmt"
	"bufio"
	"os"
)

type Client struct {
	name string
	addr string
	
}

func main(){
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
	fmt.Println("please input the client's name")
	scanner.Scan()
	name := scanner.Text()
	for  {
		fmt.Println("please input the client's address")
		scanner.Scan()
		addr := scanner.Text()

	if len(name) != 0{
		fmt.Println(name)
		fmt.Println(addr)
		arr = append(arr, name, addr)
	} else {
		break
	}
}
}
}

