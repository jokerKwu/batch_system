package ssm

import (
	"context"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var awsClientSsm *ssm.Client

func InitAws(region string) error {

	awsConfig, err := AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion(region))
	if err != nil {
		return err
	}
	awsClientSsm = ssm.NewFromConfig(awsConfig)

	return nil
}
