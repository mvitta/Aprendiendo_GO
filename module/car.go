package module

type Car struct {
	Model string `json:"Model"`
	Year  int    `json:"Year"`
	Brand string `json:"Brand"`
}

func (c Car) NewCar(model string, brand string, year int) Car {
	c.Model = model
	c.Brand = brand
	c.Year = year

	return c
}

func (c Car) GetCar() Car {
	return c
}

func (c Car) ItsNewAge() bool {
	if c.Year >= 2020 {
		return true
	}
	return false
}
