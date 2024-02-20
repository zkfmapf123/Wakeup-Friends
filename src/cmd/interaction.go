package src

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/olekukonko/tablewriter"
	src "github.com/zkfmapf123/wake-up-friends/src/aws"
)

func selectBox(msg string, list []string) string {
	prompt := &survey.Select{
		Message: msg,
		Options: list,
	}

	var answer string

	survey.AskOne(prompt, &answer, nil)
	return answer
}

func selectMultipleBox(msg string, list []string) []string {
	prompt := &survey.MultiSelect{
		Message: msg,
		Options: list,
	}

	answer := []string{}
	survey.AskOne(prompt, &answer)
	return answer
}

func PressEnter() string {
	prompt := &survey.Input{
		Message: "Press Enter the Back",
	}

	var answer string
	survey.AskOne(prompt, &answer, nil)
	return answer
}

func printTable(list map[string]src.EC2ObjectParams) {
	var tableData [][]string
	for k, v := range list {
		tableData = append(tableData, []string{k, v.Name, v.State, v.PublicIp})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"InstanceID", "Name", "State", "PublicIp"})

	for _, row := range tableData {
		table.Append(row)
	}

	table.Render()
}
