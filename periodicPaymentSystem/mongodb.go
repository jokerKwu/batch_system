package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

var PaymentCollection *mongo.Collection
var SubscriptionCollection *mongo.Collection

func InitMongoDB() error {
	clientOptions := options.Client()
	connUriDB := fmt.Sprintf("%s_%s", Env.Project, Env.Environment)
	var connUri string
	connInfos, err := AwsGetParams([]string{
		fmt.Sprintf("mongodb_%s_%s_id", Env.Environment, Env.Project),
		fmt.Sprintf("mongodb_%s_%s_pw", Env.Environment, Env.Project),
		fmt.Sprintf("mongodb_%s_domain", Env.Environment),
	})
	if err != nil {
		return err
	}
	connUriID := connInfos[0]
	connUriPW := connInfos[1]
	var connUriDomain string
	additionalOpt := ""
	if Env.IsLocal {
		connUriDomain = fmt.Sprintf("localhost:%d", 27016)
		additionalOpt = "&directConnection=true"
	} else {
		connUriDomain = connInfos[2]
	}
	if connUriID == "" || connUriPW == "" || connUriDomain == "" {
		return fmt.Errorf("no available mongodb conn info - id:%s / pw:%s / domain:%s", connUriID, connUriPW, connUriDomain)
	}
	connUri = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&replicaSet=rs0&w=majority&readPreference=primary&retryWrites=true&ssl=false%s", connUriID, connUriPW, connUriDomain, additionalOpt)

	clientOptions = clientOptions.ApplyURI(connUri)
	clientOptions.SetMaxPoolSize(1)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetMaxConnIdleTime(20 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	if err = mongoPing(); err != nil {
		return err
	}
	MongoDB = MongoClient.Database(connUriDB)
	return nil
}

func mongoPing() error {
	err := MongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}
