package pkg

import (
	"fmt"
	"log"
	"os"
)

func ReadingAndWritingFromFile() {
	file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	dataToAdd := "This is the new line to add to file"

	if _, err := file.WriteString(dataToAdd); err != nil {
		log.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		log.Println("Error while reading from file:", err)
		return
	}

	fmt.Printf("Red from file:", buffer[:n])
}