package main
import "fmt"

func main(){
	//arrays, slices, maps, loops
	//arrays ->
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr)

	//slices ->
	var intSlice []int32 = []int32{4, 5, 6}
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Println(cap(intSlice))

	var intAdd []int32 = []int32{8,9}
	intSlice = append(intSlice, intAdd...)
	fmt.Println(intSlice)

	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice2)

	// map ->
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Sabrina": 19, "Joana Paula": 20}
	fmt.Println(myMap2["Sabrina"])
	var age, its = myMap2["alguem"]
	if its{
		fmt.Printf("The age its %v", age)
	} else {
		fmt.Printf("The person doesn't exists")
	}

	for name:=range myMap2{
		fmt.Printf("nome: %v\n", name)
	}
	for name, age:=range myMap2{
		fmt.Printf("nome: %v and age %v\n", name, age)
	}

	for i, v:=range intArr{
		fmt.Printf("index: %v and value %v\n", i, v)
	}

}
