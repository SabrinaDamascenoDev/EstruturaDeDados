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
}
