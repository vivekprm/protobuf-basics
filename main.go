package main

import (
	"io/ioutil"
	"log"

	pb "github.com/vivekprm/protobuf-basics/pbgo"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	book := &pb.AddressBook{
		People: []*pb.Person{&p},
	}

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("address_book", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile("address_book")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	bookr := &pb.AddressBook{}
	if err := proto.Unmarshal(in, bookr); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	log.Printf("Address book: %+v\n", bookr)
}
