package main

import (
	"fmt"

)

type owner struct{
	name string
}
type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}
type eletricEngener struct{
	mkh uint8
	kwh uint8
}

func (e eletricEngener) milesLeft() uint8{
	return e.kwh*e.mkh
}

func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg;
}

//Interface its useful bc you can utilize a same func in multiples cases
type engine interface{
	milesLeft() uint8
}

func canMakeIt (e engine, miles uint8){
	if miles<=e.milesLeft() {
		fmt.Println("You can get in there")
	} else {
		fmt.Println("You can't get in there")
	}
}

func main(){
	var myGas gasEngine = gasEngine{23, 54, owner{"Sabrina"}}
	fmt.Printf("The mpg its: %v and have %v gallons. \n Person: %v\n", myGas.mpg, myGas.gallons, myGas.ownerInfo.name)
	fmt.Printf("The value of miles left its: %v\n", myGas.milesLeft())
	canMakeIt(myGas, 100)

	//anonymous struct
	var myGas2 = struct{
		mpg uint8
		gallons uint8
	}{23, 5}
	fmt.Println(myGas2)

	var eletric eletricEngener = eletricEngener{10, 45}
	canMakeIt(eletric, 100)
}