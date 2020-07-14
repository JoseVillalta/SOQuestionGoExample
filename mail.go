package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("test.csv")
	if err != nil{
		fmt.Errorf("Error opening File")
	}
	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err != nil{

	}
	val := record[0]

	for i := 0; i<len(val); i++{
		fmt.Printf("%x ", val[i])
	}
	name := "name"
	fmt.Println(" ")
	for i := 0; i<len(name); i++{
		fmt.Printf("%x ", name[i])
	}


	if val != "name"{
		fmt.Println("Did not match name")
	} else {
		fmt.Println("found it!")
	}

}