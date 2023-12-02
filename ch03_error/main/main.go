package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func unMarshalPerson(p Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	check(err, func() {
		log.Fatal(err)
	})
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	check(err, func() {
		log.Fatalf("error ready response body: %v", err)
	})

	err = json.Unmarshal(data, &p)
	check(err, func() {
		log.Fatalf("error unmarshalling json data: %v", err)
	})

	return
}

func check(err error, handler func()) {
	if err != nil {
		handler()
	}
}
