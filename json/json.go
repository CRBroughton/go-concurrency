package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAdrress string `json:"ip_address"`
}

func parseJSON(data []byte, channel chan Person, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	var person Person

	if err := json.Unmarshal(data, &person); err != nil {
		fmt.Println("Error; ", err)
		return
	}

	channel <- person

}

func Main() {
	file, err := os.Open("./json/MOCK_DATA.json")
	if err != nil {
		log.Fatal("Failed to open the JSON mock data file; ", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read the JSON mock data file; ", err)
	}

	var people []json.RawMessage
	if err := json.Unmarshal(bytes, &people); err != nil {
		log.Fatal("Failed to unmarshal the JSON mock data; ", err)
	}

	channel := make(chan Person)
	var waitGroup sync.WaitGroup

	for _, personData := range people {
		waitGroup.Add(1)
		go parseJSON(personData, channel, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for person := range channel {
		fmt.Printf("Parsed Person: %+v\n", person)
	}

}
