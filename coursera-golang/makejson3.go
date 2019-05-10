package main

import(
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main(){
	var(
		client map[string]string = make(map[string]string)
		scanner = bufio.NewScanner(os.Stdin)
	)
	fmt.Printf("enter client name: ")
	scanner.Scan()
	client["name"] = scanner.Text()

	fmt.Printf("Enter client address: ")
	scanner.Scan()
	client["address"] = scanner.Text()

	json, _ := json.Marshal(client)
	fmt.Printf("JSON: %v", string(json))
}


