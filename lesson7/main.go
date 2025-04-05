package main

import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32
	fmt.Println(*p)
	*p=10
	fmt.Println(*p)
	
	//p its gonna have the i memory location
	p = &i
	fmt.Println(p)
	
	
}