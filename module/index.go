package module

import (
	"fmt"
)

// Car is available in this package, doesn't require import

func Index() {
	fmt.Println("Struct and methods")
	var cars []Car
	fmt.Println(cars)

	cars = append(cars, Car{
		Model: "Renaul",
		Brand: "Logan",
		Year:  2022,
	}, Car{
		Model: "Toyota",
		Brand: "Sahara",
		Year:  2023,
	}, Car{
		Model: "3 Sedán",
		Brand: "Mazda",
		Year:  2018,
	})

	fmt.Println(cars)
	for _, value := range cars {
		car := value.GetCar()
		itsNewAge := value.ItsNewAge()
		fmt.Println(car, itsNewAge)
	}
}
