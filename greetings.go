package greetings

import (
	"errors"
	"fmt"
)

/*

This is part of "Get started" tutorial from go.dev

Just call this function to your main code and print a nice greeting

*/

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a neme was received, retun a value that embeds the name
	// in a greeting message
	message := fmt.Sprintf("✋ Hi, %v. Welcome!", name)
	return message, nil
}

/*

Note (for myself):
In Go, a function whose name starts with a capital letter
can be called by a function not in the same package

*/
