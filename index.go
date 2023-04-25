package main

import (
	"fmt"
	"math/rand"
)

func slices() {
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
	fmt.Println("suma:", sum)
	fmt.Println(list)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}

func maps() {
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

func getSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func arrays() {
	numbers := []int{
		1, 2, 3, 4, 5,
	}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("Total: ", sum)
}

func variadicFunctions(numbers ...int) {
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

func strings() {
	variable := "Hola Mundo Bello s s je je je"
	var count int
	fmt.Println(variable)
	for _, value := range variable {
		fmt.Printf("%c\n", value)
		if value == 32 {
			count++
		}
	}
	fmt.Printf("total de espacios de la franse, %q: %v", variable, count)
}

func request_api() {

}

func diferenciaEntre_Array_Slice() {
	//arrays
	var a [5]int
	array := [10]int{}

	//slices
	l := []string{}
	lista := []string{"Jean"}
	lista2 := make([]int, 3)

	fmt.Println("estos es un array", array, a)
	lista3 := append(l, "Maikol", "Hola", "Mundo")
	fmt.Println(lista)
	fmt.Println(lista2)
	fmt.Println(lista3)
	fmt.Println(l)
}

func f1() func() int {
	index := 0
	return func() int {
		index++
		return index
	}
}

func matrices() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	//remplazar pares con cero
	for row := 0; row < len(matriz); row++ {
		for colum := 0; colum < len(matriz[row]); colum++ {
			var isPar bool = matriz[row][colum]%2 == 0
			if isPar {
				matriz[row][colum] = 0
			}
		}
	}

	//sumar filas
	var sum int
	for _, value := range matriz {
		fmt.Print(value)
		sum = 0
		for i := 0; i < len(value); i++ {
			sum += value[i]
		}
		fmt.Print("-> ", sum, "\n")
	}

	//sumar columnas
	var sumCol int
	for i := 0; i < len(matriz); i++ {
		sumCol = 0
		for j := 0; j < len(matriz[i]); j++ {
			sumCol += matriz[j][i]
		}
		fmt.Print(sumCol, "  ")
	}
}

func main() {

}
