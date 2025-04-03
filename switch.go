package main;

import "fmt";

func main() {
	name := "eko"

	switch name {
	case "budi":
		fmt.Println("Halo Budi")
	case "andi":
		fmt.Println("Halo Andi")
	case "eko":
		fmt .Println("mantab Eko")
	default:
		fmt.Println("Halo, ko siapa")
	}

	switch length := len(name); length > 3 {
	case true :
		fmt.Println("nama kamu panjanggg (1)")
	case false :
		fmt.Println("mantabnyooo nama kamu (1)")
	}

	length := len(name);
	switch {
	case length > 4:
		fmt.Println("nama kamu panjanggg")
	case length < 4:
		fmt.Println("nama kamu pendek")
	default:
		fmt.Println("nama kamu sedang")
	}
}