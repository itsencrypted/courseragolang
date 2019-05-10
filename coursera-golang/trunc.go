package main
import "fmt"

var result float64

func main(){
	fmt.Println("Please enter a floating number")
	fmt.Scanln(&result)
	fmt.Printf("The truncated version of the floating point number is: ")

	fmt.Printf("%.0f",result)

}


