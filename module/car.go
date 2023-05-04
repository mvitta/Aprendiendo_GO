package module

import (
	"fmt"
	"time"
)

type Car struct {
	Model string `json:"Model"`
	Year  int    `json:"Year"`
	Brand string `json:"Brand"`
}

func (Car) GenerateCars() []Car {

	var cars []Car

	cars = append(cars, Car{
		Brand: "Reanault",
		Model: "Logan",
		Year:  2022,
	}, Car{
		Brand: "Toyota",
		Model: "Sahara",
		Year:  2023,
	}, Car{
		Brand: "Mazda",
		Model: "3 Sedán",
		Year:  2018,
	}, Car{
		Brand: "Mazda",
		Model: "3 Sedán",
		Year:  2050,
	})

	return cars
}

func (c Car) NewCar(model string, brand string, year int) Car {
	c.Brand = brand
	c.Model = model
	c.Year = year

	return c
}

func (c Car) GetCar() Car {
	return c
}

func (c Car) limit() bool {
	currentYear := time.Now().Year()
	if c.Year < currentYear {
		return true
	}
	return false
}

func (c Car) ItsNewAge() bool {
	if c.Year >= 2020 && c.limit() {
		return true
	}
	return false
}

func Show(cars []Car) {
	for i, v := range cars {
		fmt.Printf("%v	%v %v %v - It's new age: %v\n", i, v.Brand, v.Model, v.Year, v.ItsNewAge())
	}
}
