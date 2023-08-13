package main

import "fmt"

func main() {
	var (
		speed    = 100800
		distance = 96300000
	)
	const hoursPerDay = 24
	fmt.Println(distance/speed, "godzin")
	fmt.Println(distance/speed*3600, "sekund")
	fmt.Println(distance/speed/hoursPerDay, "dni")
}
