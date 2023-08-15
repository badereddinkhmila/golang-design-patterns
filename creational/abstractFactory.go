package creational

import "fmt"

/*********************************************************************/
/*************************** Interface *******************************/
/*********************************************************************/
type IMotor interface {
	setName(string)
	setPower(int)
	getName() string
	getPower() int
}

/*********************************************************************/
/*************************** Implementation **************************/
/*********************************************************************/
type Motor struct {
	name  string
	power int
}

func (m *Motor) setName(name string) {
	m.name = name
}

func (m *Motor) setPower(hpower int) {
	m.power = hpower
}

func (m *Motor) getName() string {
	return m.name
}

func (m *Motor) getPower() int {
	return m.power
}

/*********************************************************************/
/*************************** Variations ******************************/
/*********************************************************************/

type BMWMotor struct {
	Motor
}

func newBMWMotor() IMotor {
	return &BMWMotor{
		Motor: Motor{
			name:  "BMW Motor",
			power: 200,
		},
	}
}

type AudiMotor struct {
	Motor
}

func newAudiMotor() IMotor {
	return &AudiMotor{
		Motor: Motor{
			name:  "Audi Motor",
			power: 305,
		},
	}
}

type DodgeMotor struct {
	Motor
}

func newDodgeMotor() IMotor {
	return &DodgeMotor{
		Motor: Motor{
			name:  "Dodge Motor",
			power: 135,
		},
	}
}

/*********************************************************************/
/*************************** Factory Method **************************/
/*********************************************************************/
type MotorCategory string

const (
	BMWmotor   MotorCategory = "bmw_motor"
	Audimotor  MotorCategory = "audi_motor"
	Dodgemotor MotorCategory = "dodge_motor"
)

func getMotor(motorType MotorCategory) (IMotor, error) {
	if motorType == BMWmotor {
		return newBMWMotor(), nil
	} else if motorType == Dodgemotor {
		return newDodgeMotor(), nil
	} else if motorType == Audimotor {
		return newAudiMotor(), nil
	}
	return nil, fmt.Errorf("unsupported motor type passed :%s", motorType)
}

func printMotorDetails(m IMotor) {
	fmt.Printf("Motor Model: %s", m.getName())
	fmt.Println()
	fmt.Printf("Motor Power: %d hp", m.getPower())
	fmt.Println()
}

/*********************************************************************/
/*************************** Implement Abstract Factory **************/
/*********************************************************************/
type Audi struct {
}

func (a *Audi) makeCar() ICar {
	return &AudiCar{
		Car: Car{
			name:  "Audi RS8",
			power: 200,
		},
	}
}
func (a *Audi) makeMotor() IMotor {
	return &AudiMotor{
		Motor: Motor{
			name:  "Audi Motor RS9",
			power: 100,
		},
	}
}

type BMW struct {
}

func (b *BMW) makeCar() ICar {
	return &BMWCar{
		Car: Car{
			name:  "BMW M5 E60",
			power: 507,
		},
	}
}
func (b *BMW) makeMotor() IMotor {
	return &BMWMotor{
		Motor: Motor{
			name:  "BMW S 1000 RR",
			power: 205,
		},
	}
}

type Dodge struct {
}

func (d *Dodge) makeCar() ICar {
	return &DodgeCar{
		Car: Car{
			name:  "Dodge Challenger SRT Demon 170",
			power: 900,
		},
	}
}
func (d *Dodge) makeMotor() IMotor {
	return &DodgeMotor{
		Motor: Motor{
			name:  "Dodge Tomahawk",
			power: 500,
		},
	}
}

/*********************************************************************/
/*************************** Abstract Factory ************************/
/*********************************************************************/
type VehicleFamily string

const (
	BMWVehicle   VehicleFamily = "bmw"
	AudiVehicle  VehicleFamily = "audi"
	DodgeVehicle VehicleFamily = "dodge"
)

type IVehicleFactory interface {
	makeCar() ICar
	makeMotor() IMotor
}

func GetVehicleFactory(vehicleFam VehicleFamily) (IVehicleFactory, error) {
	if vehicleFam == BMWVehicle {
		return &BMW{}, nil
	}

	if vehicleFam == DodgeVehicle {
		return &Dodge{}, nil
	}

	if vehicleFam == AudiVehicle {
		return &Audi{}, nil
	}

	return nil, fmt.Errorf("wrong vehicle familt passed")
}
