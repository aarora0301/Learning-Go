package main

import "fmt"

func main(){

// declare and assign a string variable
var x string= "go"
fmt.Println(x)

// declare a string variable
var y string
y = "lang"
fmt.Println(y)

// String concatenation
fmt.Println(x+y)

// declare a string variable with double quotes
var u string= "aarshi"

// declare a string variable with back tick
var v string= `aarshi`

//compare two strings
fmt.Println("First condition: ", u==v)
fmt.Println("Second condition: ", u=="aarshi")

/* shorthand form: for string
declare and assign a variable , compiler by default infers it to string
*/
z:= "test"

fmt.Println(z)
fmt.Println("Hello World"[1])

/* shorthand form: for int
*/
o:= 5
fmt.Println(5+o)
}
