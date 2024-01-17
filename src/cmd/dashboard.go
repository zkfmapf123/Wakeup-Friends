package src

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/olekukonko/tablewriter"
	src "github.com/zkfmapf123/wake-up-friends/src/aws"
)

func GetDashboard(cfg aws.Config) string {
	list := src.GetEC2(cfg)

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

	return PressEnter()
}
