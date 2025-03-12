package main

import "fmt"

func main() {
	fmt.Println("Assalamualaikum world")

	
	var (
		firstName = "Sufardi"
		lastName = "Madoa"
	)
	fmt.Println("this my FirstName is ", firstName)
	fmt.Println("this my LastName  is ", lastName)

	const (
		nama string = "Sufardi"
		marga = "Madoa"
	)
	var name = "Sufardi"
	
	// fmt.Println("this my FirstName is constant ", nama)
	fmt.Println("this my FirstName is constant ", name)
	fmt.Println("this my FirstName is constant ", marga)

	// type Declaration
	type huruf string

	var namasaya huruf = "Ardi boy"
	var kelakuan string = "Ardi kelass"
	var ubahnama huruf = (huruf(kelakuan))
	
	fmt.Println("type data", namasaya)
	fmt.Println("type data", ubahnama)

	// operator perbandingan and or dan kebalikan
	var a = 6
	var b = 7

	var lulus bool = a > 5
	var tidaklulus bool = b > 5

	var semualulus bool = lulus && tidaklulus
	fmt.Println("jika a > 5", lulus)
	fmt.Println("jika a > 5", tidaklulus)
	fmt.Println("jika a > 5", semualulus)

	// arrayy 

	var hewan [5]string
	hewan[0] ="kambing"
	hewan[1] ="kelinci"

	fmt.Println(hewan)

	var hewan2 = [...]string{
		"kambing",
		"kelinci",
		"kelinci1",
		"kelinci2",
	}
	fmt.Println(hewan2[:])

	// make a slice
	fmt.Println("=================")
	var cobamake []string = make([]string,2,5)
	cobamake[0] = "kambing"
	cobamake[1] = "kelinci"
	cobamake2 := append(cobamake,"kelinci")
	fmt.Println(cobamake)
	fmt.Println(cobamake2)

	// make a map
	fmt.Println("=================")
	book := make(map[string]string)
	book["satu"] = "satu"

	fmt.Println(book)
}