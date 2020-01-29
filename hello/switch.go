package main

import "fmt"

func main(){

fmt.Println("Enter day")

var input int

fmt.Scan(&input)

switch input{
case 0:
case 1:
case 2:
case 3:
case 4: fmt.Println("It's a weekday")
        break
case 5:
case 6: fmt.Println("It's a weekend")
        break
default: fmt.Println("Not a valid day")
}

}
