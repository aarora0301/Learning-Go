package main

import "fmt"

func main(){

/*
create slice using make function
*/

s:= make([]string,3)
fmt.Println(s)
s[0]= "a"
s[1]="b"
s[2]="c"

fmt.Println(s)

/*
create slice using existing array
*/

arr:= [5]int{1,2,3,4,5}
// create slice x starting from index 0 to index 4  of arr
x:= arr[0:4]
fmt.Println(x)



/*
using append function
*/
slice1:= []int{1,2,3}
slice2:= append(slice1,4,5)
fmt.Println(slice1, slice2)

/*
using copy function
*/
slice3:= []int{1,2,3}
slice4:=make([]int,2)
copy(slice4, slice3)
fmt.Println(slice4)

}

