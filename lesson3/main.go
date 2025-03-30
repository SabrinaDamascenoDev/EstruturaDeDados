package main
import "fmt"
func main(){
	var StringValor string = "hihi"
	printMe()
	printMeParamens(StringValor)
}

func printMe(){
	fmt.Println("Hello World")
}

func printMeParamens(StringValue string){
	fmt.Println(StringValue)
}