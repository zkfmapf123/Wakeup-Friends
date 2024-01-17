package src

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	src "github.com/zkfmapf123/wake-up-friends/src/aws"
)

func BotherEC2(cfg aws.Config, isWakeup bool) string {

	bother, msg := getBotherSelect(isWakeup)

	list := src.GetEC2(cfg, func(status string) bool {
		return strings.Contains(bother, status)
	})

	if len(list) < 1 {
		fmt.Println("None...")
		return PressEnter()
	}

	printTable(list)
	var instanceIdAndName []string

	// instanceId/Name
	for k, v := range list {

		instanceIdAndName = append(instanceIdAndName, fmt.Sprintf("%s/%s", k, v.Name))
	}

	instand := selectBox(msg, instanceIdAndName)

	client := ec2.NewFromConfig(cfg)

	// id/name
	id := strings.Split(instand, "/")[0]

	// Running
	if isWakeup {
		input := &ec2.StartInstancesInput{
			InstanceIds: []string{id},
		}

		_, err := client.StartInstances(context.TODO(), input)
		if err != nil {
			PanicHighLight(err.Error())
		}
		// Stopped
	} else {
		input := &ec2.StopInstancesInput{
			InstanceIds: []string{id},
		}

		_, err := client.StopInstances(context.TODO(), input)
		if err != nil {
			PanicHighLight(err.Error())
		}
	}

	return PressEnter()
}

func getBotherSelect(isWakeup bool) (string, string) {
	if isWakeup {
		return "stopped", "누구를 깨울까요?"
	}

	return "running", "누구를 재울까요?"
}
