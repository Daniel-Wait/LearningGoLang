package greetings

import "errors"
import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

    // Return a greeting that embeds the name in a message.
	if false {
		message := fmt.Sprintf("Hi, %v. Welcome!", name)
		return message, nil
	} else {
		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
		return message, nil
	}
}