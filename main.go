package main //this line is in every go progrom

import "fmt" // packages

func main() { 
	var name string //varibles in go

	fmt.Print("enter your name : ")
	fmt.Scan(&name)//scans for use input  
	fmt.Println("hello", name , "nice to meet you")
}
// this is how you scan for user input in varibles 


