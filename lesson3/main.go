package main
import (
	"fmt"
	"errors"
)
func main(){
	var StringValor string = "hihi"
	printMe()
	printMeParamens(StringValor)
	fmt.Println(division(10, 2))

	var div, rest, err = divisionAndRest(10, 2)
	//Structure switch with errors
	switch{
		case err!=nil:
			fmt.Println(err.Error())
		case rest == 0:
			fmt.Printf("The result its %v with no rest", div)
		default:
			fmt.Printf("the result its %v and the rest its %v", div, rest)
	}
	// if else with erros
	if (err!=nil) {
		fmt.Println(err.Error())
	} else if(rest == 0){
		fmt.Printf("The result its %v with no rest", div)
	} else{
		fmt.Printf("the result its %v and the rest its %v", div, rest)
	}
	

}

func printMe(){
	fmt.Println("Hello World")
}

func printMeParamens(StringValue string){
	fmt.Println(StringValue)
}

func division(numerador int, denominador int) int{
	var resul int = numerador/denominador
	return resul
}

func divisionAndRest(numerador int, denominador int) (int, int, error){
	var err error
	if denominador==0{
		err = errors.New("The donominator it's cannot be 0")
		return 0, 0, err
	}
	var resul int = numerador/denominador
	var rest int = numerador%denominador
	return resul, rest, err
}