package main

import "fmt"

func main(){

/*
create map
*/
m:= make(map[string]int)

/*
add key value pairs 
*/
m["first"]= 1
m["second"]=2

fmt.Println(m)

/*
delete key from map
*/
delete(m, "first")
fmt.Println(m)

/*
Access a key
*/

fmt.Println(m["second"])

/*
creating maps of maps
*/
elements:= map[string]map[string]string{

"He": map[string]string{
      "name": "helium",
      "state":"gas",
},
"H": map[string]string{
"name":"hydrogen",
"state":"gas",
},
}
fmt.Println(elements["H"])

}
