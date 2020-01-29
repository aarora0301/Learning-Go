package main
import "fmt"

func main(){

var x[5]int
x[4]=100

fmt.Println(x)
test()
test1()
}

func test(){
nums:= [] int{2,3,4}
/*
Iterates through array nums, 
num is the num in array
_ is is the index, _ is used when index is not required in array iteration
*/
for _, num := range nums{
fmt.Println(num)
}
}

func test1(){
var total=0
nums:=[] int{2,3,4}
for i, num:=range nums{
total+=num
fmt.Println(i)
fmt.Println(total)
}


}
