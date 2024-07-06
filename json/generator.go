package json

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var firstNames = []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie"}
var lastNames = []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis"}
var genders = []string{"Male", "Female", "Non-binary", "Other"}
var domains = []string{"example.com", "mail.com", "test.com"}

func randomStringFromSlice(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func randomEmail(firstName, lastName string) string {
	return fmt.Sprintf("%s.%s@%s", firstName, lastName, randomStringFromSlice(domains))
}

func randomIPAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func Generator() {
	rand.Seed(time.Now().UnixNano())

	n := 50000
	people := make([]Person, n)

	for i := 0; i < n; i++ {
		firstName := randomStringFromSlice(firstNames)
		lastName := randomStringFromSlice(lastNames)
		people[i] = Person{
			Id:        i + 1,
			FirstName: firstName,
			LastName:  lastName,
			Email:     randomEmail(firstName, lastName),
			Gender:    randomStringFromSlice(genders),
			IPAddress: randomIPAddress(),
		}
	}

	file, err := os.Create("large_data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(people); err != nil {
		fmt.Println("Error:", err)
	}
}
