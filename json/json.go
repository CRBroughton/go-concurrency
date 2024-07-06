package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
}

func parseJSONSequential(data []json.RawMessage) []Person {
	var people []Person
	for _, personData := range data {
		var person Person
		if err := json.Unmarshal(personData, &person); err != nil {
			log.Fatal("Failed to unmarshal single person; ", err)
		}
		people = append(people, person)
	}
	return people
}

func parseJSON(data []byte, channel chan<- Person, wg *sync.WaitGroup) {
	defer wg.Done()
	var person Person
	if err := json.Unmarshal(data, &person); err != nil {
		fmt.Println("Error:", err)
		return
	}
	channel <- person
}

func parseJSONConcurrently(peopleData []json.RawMessage) []Person {
	channel := make(chan Person, len(peopleData))
	var waitGroup sync.WaitGroup

	for _, personData := range peopleData {
		waitGroup.Add(1)
		go parseJSON(personData, channel, &waitGroup)
	}

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	var people []Person
	for person := range channel {
		people = append(people, person)
	}
	return people
}

func JSONMain() {
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

	// Sequential Execution
	startSequential := time.Now()
	peopleSequential := parseJSONSequential(people)
	durationSequential := time.Since(startSequential)

	fmt.Printf("Sequential Parsed People: %d\n", len(peopleSequential))
	fmt.Printf("Sequential Execution Time: %v\n", durationSequential)

	// Concurrent Execution
	startConcurrent := time.Now()
	peopleConcurrent := parseJSONConcurrently(people)
	durationConcurrent := time.Since(startConcurrent)

	fmt.Printf("Concurrent Parsed People: %d\n", len(peopleConcurrent))
	fmt.Printf("Concurrent Execution Time: %v\n", durationConcurrent)
}
