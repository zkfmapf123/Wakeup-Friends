package cmd

import "github.com/AlecAivazis/survey/v2"

func selectBox(msg string, list []string) string {
	prompt := &survey.Select{
		Message: msg,
		Options: list,
	}

	var answer string

	survey.AskOne(prompt, &answer, nil)
	return answer
}
