package main
import "fmt"

var fahrenheit float64
var celsius float64

func main(){
fmt.Print("Enter the temperature (Fahrenheits): ")
fmt.Scanln(&fahrenheit)

celsius := (fahrenheit - 32) * (5.0/9)
fmt.Println(celsius, "C")
}


