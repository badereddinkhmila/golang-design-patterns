package creational

import (
	"fmt"
	"time"
)

func CreationalDP() {
	fmt.Println("<----- Creational Start ----->")
	fmt.Println("\n<----- Singleton ----->")
	for i := 0; i < 5; i++ {
		go getInstance()
		go getInstanceOnce()
	}
	time.Sleep(time.Duration(50 * time.Millisecond))

	fmt.Println("\n<----- Factory Method ----->")
	bmw, err := getCar(CarCategory("bmw"))
	if err != nil {
		fmt.Println(err)
	}
	printCarDetails(bmw)
	dodge, err := getCar(CarCategory("dodge"))
	if err != nil {
		fmt.Println(err)
	}
	printCarDetails(dodge)

	fmt.Println("\n<----- Abstract Factory ----->")
	bmwFactory, err := GetVehicleFactory(BMWVehicle)
	if err != nil {
		fmt.Println(err)
	}
	bmwCar := bmwFactory.makeCar()
	bmwMotor := bmwFactory.makeMotor()
	printCarDetails(bmwCar)
	printMotorDetails(bmwMotor)
	fmt.Println("\n<----- Prototype ----->")
	s1 := &Student{Firstname: "Badreddine", Age: 27}
	fmt.Printf("[Simple obj] %s\n", s1.Firstname)
	s1Copy := s1.clone()
	s1Copy.print("[Simple obj copy] ")
	s2 := &Student{Firstname: "Ayoub", Age: 13}
	s3 := &Student{Firstname: "Yassine", Age: 23}
	class := &Classroom{Floor: "3rd", Tag: "Lab", students: []IClonable{s1, s2, s3}}
	classClone := class.clone()
	classClone.print("[Complex obj copy] ")

	fmt.Println("\n<----- Builder ----->")
	normalBuilder, _ := getHouseBuilder(Normal)
	iglooBuilder, _ := getHouseBuilder(IGloo)

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorsType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowsType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floors)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorsType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowsType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floors)
	fmt.Println("\n<----- Builder Functional ----->")
	address := &Address{City: "Sfax", Postcode: 3040, Country: "Tunisia"}

	user1 := NewUser(nil, nil, "1234", "test@go.world", nil, address)
	fmt.Printf("[All fields builder] %v", user1)

	user2 := InitUser("test@go.world", "1234").setFirstname("badreddine").setLastname("khmila").setAge(27).setAddress(address)
	fmt.Printf("[Field by field builder] %v", user2)

	fmt.Println("\n<----- Creational Done ----->")
}
