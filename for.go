package main

import "fmt"

func main(){
	
	// basic for
	for i := 0; i < 10; i++{
		fmt.Println("maka hasil yang kita dapat ", i)
	
	}
	// cara pertama
	names := []string{"eko", "budi", "Sufardi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// cara kedua
	for index := range names {

		// fmt.Println("index", index, "namanya", name)
		if index%2 == 1 {
			fmt.Println("ini adalah",index)
			continue
		}
	}
}