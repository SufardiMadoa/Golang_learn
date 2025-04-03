package main

import "fmt"

func Assalamualaikum() {
	fmt.Println("Assalamualaikum")
}

func mantab( name string, umur int){
	fmt.Println("Halo", name, "umur kamu", umur)

}
func returnvalue(name string) string {
	return "hello " + name
}
func returnmultiple() (string, int){
	return "Sufardi", 22
}
func namedreturn()(name, age, mantab string) {
	name = "Sufardi"
	age = "22"
	mantab = "mantab"
	return name, age, mantab
}
func main(){
	// Assalamualaikum()
	// mantab("sufardi", 22)
	// result := returnvalue("Sufardi")
	// fmt.Println(result)
	
	// nama,_ := returnmultiple()
	// fmt.Println("nama saya",nama)

	nama, umur, mantab :=namedreturn()
	fmt.Println("nama saya",nama,"umur saya",umur,mantab)

}