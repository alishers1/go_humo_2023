package pkg

import (
	"encoding/json"
	"log"
	"os"
)

type Student struct {
	FirstName   string
	LastName    string
	Age         uint
	PhoneNumber string
	IsMarried   bool
}

func WriteInFile() {
	student := Student{
		FirstName:   "Alisher",
		LastName:    "Siddiqov",
		Age:         23,
		PhoneNumber: "555577709",
		IsMarried:   false,
	}

	file, err := os.Create("data.json")
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	data, err := json.MarshalIndent(student, "", "       ")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
}
