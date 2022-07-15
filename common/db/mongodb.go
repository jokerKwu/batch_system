package db

import (
	"batch_system/common/aws/ssm"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func (a *EnvMongoDB) MakeConnURI(isLocal bool) string {
	var connUri string

	connInfos, _ := a.GetSsmMongoInfo()
	connUriID := connInfos[0]
	connUriPW := connInfos[1]
	additionalOpt := ""
	connUriDomain := connInfos[2]
	connUri = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&replicaSet=%s&w=majority&readPreference=primary&retryWrites=true&ssl=false%s", connUriID, connUriPW, connUriDomain, a.RsName, additionalOpt)
	if isLocal == true {
		connUriDomain = fmt.Sprintf("localhost:%d", a.LocalPort)
		additionalOpt = "&directConnection=true"
		connUri = fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&replicaSet=%s&w=majority&readPreference=primary&retryWrites=true&ssl=false%s", connUriID, connUriPW, connUriDomain, a.RsName, additionalOpt)
	}

	return connUri
}

func (a *EnvMongoDB) ConnectMongo(connUri string) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI(connUri)
	clientOptions.SetMaxPoolSize(200)
	clientOptions.SetMinPoolSize(10)
	clientOptions.SetMaxConnIdleTime(10 * time.Second)
	clientOptions.SetSocketTimeout(8 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return &mongo.Client{}, err
	}
	return mongoClient, nil
}

func (a *EnvMongoDB) InitCollection() error {

	return nil
}

func (a *EnvMongoDB) PingMongo(mongoClient *mongo.Client) error {
	err := mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func (a *EnvMongoDB) GetSsmMongoInfo() ([]string, error) {
	var connInfos []string

	connInfos, err := ssm.AwsGetParams([]string{
		fmt.Sprintf("mongodb_%s_%s_id", a.Env, a.Project),
		fmt.Sprintf("mongodb_%s_%s_pw", a.Env, a.Project),
		fmt.Sprintf("mongodb_%s_%s_domain", a.Env, a.Project),
	})
	if err != nil {
		return nil, err
	}

	return connInfos, nil
}

func CloseMongo() {
	_ = MongoClient.Disconnect(context.TODO())
}
