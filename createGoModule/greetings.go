package main

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.package createGoModule
	message := fmt.Sprintf("Hi, ½v. Welcome!", name)
	return message
}