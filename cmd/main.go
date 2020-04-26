package main

import (
	"fmt"
	"log"
	"main/protobuf_data"

	"github.com/golang/protobuf/proto"
)

func main() {

	personTom := &protobuf_data.Person{
		Name: "Tom",
		Age:  32,
	}

	data := marshaling(personTom)
	unMarshalling(data)
}

func marshaling(person *protobuf_data.Person) []byte {

	data, err := proto.Marshal(person)

	if err != nil {
		log.Fatal("Marshling Error: ", err)
	}

	fmt.Println(data) // Printing out raw protobuf object
	return data

}

// unmarshal byte array into an object we can modify and use
func unMarshalling(data []byte) {

	newPersonTom := &protobuf_data.Person{}
	err := proto.Unmarshal(data, newPersonTom)

	if err != nil {
		log.Fatal("Unmarshling Error: ", err)
	}

	fmt.Println(newPersonTom.Name)
	fmt.Println(newPersonTom.Age)

}
