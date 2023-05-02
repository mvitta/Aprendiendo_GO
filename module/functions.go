package module

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reflect"
)

func Slices() {
	list := make([]int, 10)
	for i := 0; i < len(list); i++ {
		number := rand.Intn(20)
		list[i] = number
	}
	elementSum := list[5:]
	fmt.Println(elementSum)
	sum := 0
	for j := 0; j < len(elementSum); j++ {
		sum += elementSum[j]
	}
	fmt.Println("sum:", sum)
	fmt.Println(list)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}

func Maps() {
	myMap := make(map[string]int)
	myMap["keyOne"] = 10
	myMap["keyTwo"] = 20
	fmt.Println(myMap)

	//------------------------------
	otherMap := map[string]int{"k1": 20, "k2": 30}
	fmt.Println(otherMap)

	//---------------------------------

	newMyMap := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range newMyMap {
		fmt.Println(key, value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	//------------------------------------

}

func GetSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func Arrays() {
	numbers := []int{
		1, 2, 3, 4, 5,
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Total: ", sum)
}

func VariadicFunctions(numbers ...int) {
	fmt.Println(numbers)
	fmt.Printf("%T\n", numbers)
	sum := 0
	for index, value := range numbers {
		if index%2 == 0 {
			sum += value
			fmt.Println("index: ", index, sum)
		}
	}
	fmt.Println("\ntotal: ", sum)
}

func Strings() {
	variable := "Hello beautiful Word s s je je je"
	var count int
	fmt.Println(variable)
	for _, value := range variable {
		fmt.Printf("%c\n", value)
		if value == 32 {
			count++
		}
	}
	fmt.Printf("total spaces of the phrase, %q: %v", variable, count)
}

func Request_api() {

}

func Array_Slice() {
	//arrays
	var a [5]int
	array := [10]int{}

	//slices
	l := []string{}
	list := []string{"Jean"}
	list2 := make([]int, 3)

	fmt.Println("this is an array", array, a)
	list3 := append(l, "Michael", "Hello", "Word")
	fmt.Println(list)
	fmt.Println(list2)
	fmt.Println(list3)
	fmt.Println(l)
}

func F1() func() int {
	index := 0
	return func() int {
		index++
		return index
	}
}

func Matrices() {
	var m [10][19]int
	fmt.Println(m)
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//replace even number with zero
	for row := 0; row < len(matrix); row++ {
		for colum := 0; colum < len(matrix[row]); colum++ {
			var isPar bool = matrix[row][colum]%2 == 0
			if isPar {
				matrix[row][colum] = 0
			}
		}
	}

	//add columns
	var sum int
	for _, value := range matrix {
		fmt.Print(value)
		sum = 0
		for i := 0; i < len(value); i++ {
			sum += value[i]
		}
		fmt.Print("-> ", sum, "\n")
	}

	//add columns
	var sumCol int
	for i := 0; i < len(matrix); i++ {
		sumCol = 0
		for j := 0; j < len(matrix[i]); j++ {
			sumCol += matrix[j][i]
		}
		fmt.Print(sumCol, "  ")
	}
}

func Pointers() {
	var pointer *int
	array := [3]int{
		10, 20, 30,
	}
	indexRandom := rand.Intn(3)
	fmt.Println(indexRandom)
	for i := 0; i < len(array); i++ {
		fmt.Println("value:", array[i], "memory:", &array[i], ".")
	}
	pointer = &array[indexRandom]
	fmt.Println(*pointer)
	*pointer = 300
	fmt.Println(array)
}

// not optimal, start traversing from the ends.
// just to understand how pointers work
func Pointers2() (int, int) {
	var pointer *string
	myArray := []string{"a", "b", "c", "b", "z"}
	alternateIndex := len(myArray) - 1
	for i := 0; i < len(myArray); i++ {
		pointer = &myArray[alternateIndex]
		fmt.Println(myArray[i], *pointer)
		if myArray[i] == *pointer {
			return i, alternateIndex
		}
		alternateIndex--
	}
	return -1, -1
}

func Average(arr *[6]float64, sum float64) float64 {
	var average float64
	for _, value := range *arr {
		sum += value
	}
	average = sum / float64(len(arr))
	fmt.Println(average)
	return 0
}

func DeleteElementSlice() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
}

type Person struct {
	name string
	age  int
}

func (p Person) Structs(number int) {
	persona := Person{}
	fmt.Println(p)
	fmt.Println(persona)
	fmt.Println(number)

}

func TojsonStruct() {
	// yo can create a type that accepts multiple types
	type ColorGroups struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroups{
		ID:     1,
		Name:   "Red",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	//here... type ColorGroups
	fmt.Println("Type ColorGroups", group)

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error", err)
	}
	//here it is converted to json
	os.Stdout.Write(b)
	//___________________________________________________________
	getGroup := ColorGroups{}
	//here the json will be converted to type ColorGroups
	err1 := json.Unmarshal(b, &getGroup)
	if err != nil {
		fmt.Println("Error", err1)
	}
	fmt.Println()

	fmt.Printf("%+v", getGroup)

}

func TojsonMap() {
	// all values must be of the same type
	dataMap := map[string]string{
		"firstName": "Michael",
		"lastName":  "Vitta",
		"Cedula":    "1234567890",
	}

	b, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("%s", b)
}

func Marshal() {
	type persons struct {
		FirstName string
		LastName  string
	}

	listPerson := []persons{
		{FirstName: "Michael", LastName: "Vitta"},
		{FirstName: "Jean", LastName: "Vitta"},
		{FirstName: "Sheryl", LastName: "Vitta"},
	}
	fmt.Println(listPerson)
	fmt.Println()

	bytes, err := json.Marshal(listPerson)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(bytes))
}

func Unmarshal() {
	type person struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
	}

	sliceBytes := []byte(`[{"FirstName":"Michael","LastName":"Vitta"},{"FirstName":"Jean","LastName":"Vitta"},{"FirstName":"Sheryl","LastName":"Vitta"}]`)

	var p []person
	fmt.Println(p)
	err := json.Unmarshal(sliceBytes, &p)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(p)
	fmt.Println(reflect.TypeOf(p))
}
