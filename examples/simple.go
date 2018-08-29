package main

import (
	"fmt"

	"github.com/tylerflint/survey"
)

// the questions to ask
var simpleQs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "What is your name?",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
		},
		Validate: survey.Required,
	},
	{
		Name: "password",
		Prompt: &survey.Password{
			Message: "Password:",
		},
	},
}

func main() {
	answers := struct {
		Name  		string
		Color 		string
		Password 	string
	}{}

	// ask the question
	err := survey.Ask(simpleQs, &answers)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// print the answers
	fmt.Printf("%s chose %s with password: %s.\n", answers.Name, answers.Color, answers.Password)
}
