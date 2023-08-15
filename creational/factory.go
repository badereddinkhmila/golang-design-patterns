package creational

import "fmt"

/*********************************************************************/
/*************************** Interface *******************************/
/*********************************************************************/
type ICar interface {
	setName(string)
	setPower(int)
	getName() string
	getPower() int
}

/*********************************************************************/
/*************************** Implementation **************************/
/*********************************************************************/
type Car struct {
	name  string
	power int
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) setPower(hpower int) {
	c.power = hpower
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) getPower() int {
	return c.power
}

/*********************************************************************/
/*************************** Variations ******************************/
/*********************************************************************/

type BMWCar struct {
	Car
}

func newBMW() ICar {
	return &BMWCar{
		Car: Car{
			name:  "BMW m5",
			power: 400,
		},
	}
}

type AudiCar struct {
	Car
}

func newAudi() ICar {
	return &AudiCar{
		Car: Car{
			name:  "Audi Q8",
			power: 30,
		},
	}
}

type DodgeCar struct {
	Car
}

func newDodge() ICar {
	return &DodgeCar{
		Car: Car{
			name:  "Dodge challenger",
			power: 350,
		},
	}
}

/*********************************************************************/
/*************************** Factory Method **************************/
/*********************************************************************/
type CarCategory string

const (
	BMWModel   CarCategory = "bmw"
	AudiModel  CarCategory = "audi"
	DodgeModel CarCategory = "dodge"
)

func getCar(carType CarCategory) (ICar, error) {
	if carType == BMWModel {
		return newBMW(), nil
	} else if carType == DodgeModel {
		return newDodge(), nil
	} else if carType == AudiModel {
		return newAudi(), nil
	}
	return nil, fmt.Errorf("unsupported car type passed :%s", carType)
}

func printCarDetails(c ICar) {
	fmt.Printf("Car Model: %s", c.getName())
	fmt.Println()
	fmt.Printf("Car Power: %d hp", c.getPower())
	fmt.Println()
}
