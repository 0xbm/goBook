package main

import "fmt"

func main() {
	const lightspeed = 299792
	var distance = 56000000
	fmt.Println(distance/lightspeed, "sekund")
	distance = 401000000
	fmt.Println(distance/lightspeed, "sekund")
}
