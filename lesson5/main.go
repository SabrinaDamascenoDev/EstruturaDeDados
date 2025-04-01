package main
import (
	"fmt"
	"strings"
)
func main(){
	var myString = "Résumé"
	var theValueOfTheFirstChar = myString[0]
	fmt.Printf("%v, %T\n", theValueOfTheFirstChar, theValueOfTheFirstChar)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	//Mais devagar para criar uma string
	var nome = []string{"s", "a", "b", "r", "i", "n", "a"}
	var nomeMostrar = ""
	for i := range nome{
		nomeMostrar += nome[i]
	}
	fmt.Println(nomeMostrar)

	//Forma mais rápida, com o strings
	var stgBuilder strings.Builder
	for i := range nome{
		stgBuilder.WriteString(nome[i])
	}
	var catStg = stgBuilder.String()
	fmt.Println(catStg)
}