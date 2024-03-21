package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "DOG"
}

type Robot struct{}

func (r *Robot) Speak() string {
	return "ROBOT"
}

func introSpeaker(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func main() {
	d := &Dog{}
	r := &Robot{}

	introSpeaker(d)
	introSpeaker(r)
}
