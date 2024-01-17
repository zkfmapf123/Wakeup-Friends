package src

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2ObjectParams struct {
	State    string
	PublicIp string
	Name     string
}

func GetEC2(cfg aws.Config) map[string]EC2ObjectParams {
	client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstancesInput{}

	res, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		// pani
	}

	ec2Object := make(map[string]EC2ObjectParams)
	for _, reservation := range res.Reservations {
		for _, ins := range reservation.Instances {
			insId, insState, publicIp := ins.InstanceId, ins.State.Name, ins.PublicIpAddress

			for _, tag := range ins.Tags {
				if *tag.Key == "Name" {

					ec2Object[*insId] = EC2ObjectParams{
						State:    string(insState),
						PublicIp: getPublicInstance(publicIp),
						Name:     *tag.Value,
					}
				}
			}
		}
	}

	return ec2Object
}

func getPublicInstance(ip *string) string {
	if ip != nil {
		return *ip
	}

	return "none"
}
