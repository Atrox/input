package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ValidatorFunction describes the simple interface every validator needs to meet
type ValidatorFunction func(string) (interface{}, error)

// Prompt prints the question to Stdout and checks the user input according to the provided validators
func Prompt(question string, validators ...ValidatorFunction) (out interface{}) {
	for {
		fmt.Printf("[?] %s: ", question)

		input, err := getInput()
		if err != nil {
			log.Fatal(err)
		}

		var validatorError error
		for _, validator := range validators {
			out, err = validator(input)
			if err != nil {
				validatorError = err
				break
			}
		}

		if validatorError == nil {
			break
		}

		fmt.Printf("[!] Error: %s\n", validatorError.Error())
	}
	return
}

// Create a single reader which can be called multiple times
var reader = bufio.NewReader(os.Stdin)

// GetInput gets the input from os.Stdin and trims it
func getInput() (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
