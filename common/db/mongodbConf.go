package db

import "go.mongodb.org/mongo-driver/mongo"

type EnvMongoDB struct {
	Project   string
	Env       string
	RsName    string
	LocalPort int64
}

type MongoConf interface {
	GetSsmMongoInfo() ([]string, error)
	ConnectMongo(connInfos []string, isLocal bool) (*mongo.Client, error)
	InitCollection() error
	PingMongo(mongoClient *mongo.Client) error
}

var (
	UserCollection             *mongo.Collection
	SubscriptionCollection     *mongo.Collection
	SubscriptionPlanCollection *mongo.Collection
	AddressBookCollection      *mongo.Collection
	DeliveryCollection         *mongo.Collection
	TermsCollection            *mongo.Collection
	ProductInfoCollection      *mongo.Collection
	QnACollection              *mongo.Collection
	PaymentCollection          *mongo.Collection
	EmailAuthCollection        *mongo.Collection
	DeliveryHistoryCollection  *mongo.Collection
	PaymentHistoryCollection   *mongo.Collection

	AppUserAuthCollection *mongo.Collection
	AppUserCollection     *mongo.Collection
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database
