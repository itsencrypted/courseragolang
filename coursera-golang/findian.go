 package main
 import( "fmt"
 "strings"
 )

 func main(){
	 var name string
	 fmt.Print("Please enter a string of your choice: ")
	 fmt.Scanln(&name)
	 name = strings.ToLower(name)

	  if strings.HasPrefix(name, "i")&&
	  strings.HasSuffix(name,"n")&&
	  strings.ContainsAny(name, "a"){
		  fmt.Println("Found!")
	  }else{
		  fmt.Println("Not Found!")
	  }
  }



