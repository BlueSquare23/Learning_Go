package main

import "fmt"

// 1. Define the interface
type Speaker interface {
	Speak() string
}

// 2. Define a Dog struct and its Speak() method
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

// 3. Define a Robot struct and its Speak() method
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep boop! Model " + r.Model + " operational."
}

// 4. A function that accepts ANY Speaker (Polymorphism)
func MakeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	myDog := Dog{Name: "Buddy"}
	myRobot := Robot{Model: "T-800"}

	// Both work perfectly because they both have a Speak() method!
	MakeItSpeak(myDog)   // Output: Woof! My name is Buddy
	MakeItSpeak(myRobot) // Output: Beep boop! Model T-800 operational.
}
