package tools

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"
)

type MongoDBSetting struct {
	TimeoutSecs int    `mapstructure:"MongoTimeoutSecs"`
	IP          string `mapstructure:"MongoIP"`
	PORT        int    `mapstructure:"MongoPORT"`
}

type MongoDBManager struct {
	Setting *MongoDBSetting
	Client  *mongo.Client
}

func (Manager *MongoDBManager) Init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Manager.Setting.TimeoutSecs)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://"+Manager.Setting.IP+":"+strconv.Itoa(Manager.Setting.PORT)))
	if err != nil {
		panic(err)
	}
	Manager.Client = client
}

func (Manager *MongoDBManager) IsConnected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Manager.Setting.TimeoutSecs)*time.Second)
	defer cancel()
	err := Manager.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (Manager *MongoDBManager) ProvideDBConnection() any {
	return Manager.Client
}

func NewMongoDBManager(Setting *MongoDBSetting) *MongoDBManager {
	Manager := MongoDBManager{Setting: Setting}
	Manager.Init()
	return &Manager
}
