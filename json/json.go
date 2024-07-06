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
	IPAdrress string `json:"ip_address"`
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

func parseJSON(data []byte, channel chan Person, wg *sync.WaitGroup) {
	defer wg.Done()
	var person Person
	if err := json.Unmarshal(data, &person); err != nil {
		fmt.Println("Error:", err)
		return
	}
	channel <- person
}
func parseJSONConcurrently(peopleData []json.RawMessage) []Person {
	c := make(chan Person)
	var wg sync.WaitGroup

	for _, personData := range peopleData {
		wg.Add(1)
		go parseJSON(personData, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var people []Person
	for person := range c {
		people = append(people, person)
	}
	return people
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

	// Sequential reading
	start := time.Now()
	sequentialPeople := parseJSONSequential(people)
	duration := time.Since(start)

	fmt.Printf("Parsed Sequential People: %+v\n", sequentialPeople)
	fmt.Printf("Sequential Execution Time: %v\n", duration)

	// // Concurrent Execution
	// startConcurrent := time.Now()
	// peopleConcurrent := parseJSONConcurrently(people)
	// durationConcurrent := time.Since(startConcurrent)

	// fmt.Printf("Concurrent Parsed People: %+v\n", peopleConcurrent)
	// fmt.Printf("Concurrent Execution Time: %v\n", durationConcurrent)

}
