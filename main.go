package main

import (
	"fmt"
	"log"
	"example.com/server/docum"
)

func main() {
	fmt.Println("hellow go from muscular memory");
	// lse.Pointers();

	names := []string{"Gwyne","Roxx","Juan"}
	messages, err := docum.Hellos(names)
	if err!=nil{
		log.Fatal("Error doing hellos")
	}

	fmt.Println(messages)


}
