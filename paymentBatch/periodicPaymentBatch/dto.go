package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SubscriptionDTO struct {
	ID                 primitive.ObjectID    `bson:"_id,omitempty"`
	SubscriptionPlanID primitive.ObjectID    `bson:"subscriptionPlanID,omitempty"`
	AddressBookID      primitive.ObjectID    `bson:"addressBookID,omitempty"`
	UserID             primitive.ObjectID    `bson:"userID,omitempty"`
	SubscriptionType   SubscriptionType      `bson:"subscriptionType"`
	MemberCount        int64                 `bson:"memberCount"`
	State              SubscriptionState     `bson:"state"`
	Created            time.Time             `bson:"created,omitempty"`
	StartDate          time.Time             `bson:"startDate,omitempty"`
	IsDeleted          bool                  `bson:"isDeleted"`
	LastUpdate         time.Time             `bson:"lastUpdate,omitempty"`
	ProductList        []SubscriptionProduct `bson:"productList,omitempty"`
	NextPayment        *time.Time            `bson:"nextPayment,omitempty"`
}
type SubscriptionProduct struct {
	ProductID primitive.ObjectID `bson:"productID,omitempty"`
	Amount    int64              `bson:"amount,omitempty"`
}

type SubscriptionState int64

const (
	SUBS_READY  = SubscriptionState(1)
	SUBS_WAIT   = SubscriptionState(2)
	SUBSCRIBING = SubscriptionState(3)
	SUBS_CANCEL = SubscriptionState(4)
)

type SubscriptionType int64

const (
	BASIC = SubscriptionType(1)
	PLUS  = SubscriptionType(2)
)
