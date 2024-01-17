package src

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	src "github.com/zkfmapf123/wake-up-friends/src/aws"
)

func GetDashboard(cfg aws.Config) string {
	list := src.GetEC2(cfg, func(status string) bool {
		allState := []string{"running", "stopped", "stopping", "pending", "shutting-down", "terminated"}
		for _, v := range allState {
			if strings.Contains(v, status) {
				return true
			}
		}

		return false
	})

	printTable(list)
	return PressEnter()
}
