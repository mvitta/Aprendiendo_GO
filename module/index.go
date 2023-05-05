package module

import (
	"fmt"
	"sort"
)

// Car is available in this package, doesn't require import

// second way -> sort.Sort() ...
type ByYear []Car

func (y ByYear) Len() int {
	return len(y)
}

func (y ByYear) Swap(current, next int) { y[current], y[next] = y[next], y[current] }

func (y ByYear) Less(current, next int) bool {
	return y[current].Year > y[next].Year
}

func Index() {
	fmt.Println("Struct and methods")
	cars := Car{}.GenerateCars()

	Show(cars)
	cars[3].SetCar("RX-7", "Mazda", 3000)
	//first form
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Year < cars[j].Year
	})

	// ... second way -> sort.Sort()

	sort.Sort(ByYear(cars))
	fmt.Println()
	Show(cars)

}

// func (p Person) String() string {
// 	return fmt.Sprintf("%s: %d", p.Name, p.Age)
// }
