 package main

 import "fmt"

 func main() {
	 var f float64
	 fmt.Println("Please enter a float")
	 _, err := fmt.Scan(&f)
	 if err !=nil {
		 fmt.Println("Input was not a float.Please re-run the program and input a float next time")
		 return
	 }

	 fmt.Printf("%d", int(f))
	 fmt.Println()
 }


