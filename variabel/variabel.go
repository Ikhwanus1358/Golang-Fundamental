package main

import "fmt"

func main() {
	var firstName string = "Ikhwanushofa"
	var lastName string = "Effendi"

	fmt.Println(firstName, lastName)
	fmt.Printf("Hallo, Ikhwanushofa Effendi\n")
	fmt.Printf("Hallo, %s %s \n", firstName, lastName)

	var (
		age     int
		address string
	)

	age = 23
	address = "Bandung"

	fmt.Printf("Hallo, saya %s %s berumur %d dan saya tinggal di %s\n", firstName, lastName, age, address)

	var (
		bootcampName, bootcampAddress = "Waanus Camp", "Bandung"
	)

	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := "31"
	month := "October"

	fmt.Println(day, date, month+" 2023")

	var number = 21

	number = 20

	const phi = 3.14

	fmt.Println(number, phi)

}
