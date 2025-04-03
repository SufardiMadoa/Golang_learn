package main;

import "fmt";

func main() {
	name := "budi"

	if name == "eko" {
		fmt.Println("Hello, eko!")
	} else if name == "budi" {
	fmt.Println("mantab Budi")
	} else {
		fmt.Println("daa ko")
	}
	if length := len(name); length > 3 {
		fmt.Println("nama kamu panjang")
	}
}