package main

import "fmt"

func main(){
	//array1 :=[5]string{"i", "n", "h", "a", "!"}
	array3 := [5]string{"a", "b", "c", "d", "e"}
	array3[2] = "2"
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	//slice1 := array1[2:]
	fmt.Println(slice3, slice4)
	slice3[2] = "Queen"
	fmt.Println(slice3, slice4)
	fmt.Println(array3)
	//fmt.Println(slice1)
}
