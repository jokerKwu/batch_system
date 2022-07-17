package main

import (
	"fmt"
	common "github.com/jokerkwu/backend_common/batch_common"
	"go.mongodb.org/mongo-driver/mongo"
)

var SubscriptionCollection *mongo.Collection

func InitMongoDB() error {
	var err error
	err = common.InitMongoDB()
	if err != nil {
		fmt.Println("mongo client init")
		return err
	}

	SubscriptionCollection = common.MongoDB.Collection("subscription")

	return nil
}
