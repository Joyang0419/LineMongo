package userMessageRepo

import (
	"LineMongo/internals/models"
	"LineMongo/internals/tools"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUserMessageRepo struct {
	MongoDBManager *tools.MongoDBManager
}

func (repo *MongoUserMessageRepo) Get(filters ...Filter) (UserMessageDTOs []map[string]any) {
	collection := repo.MongoDBManager.Client.Database("test").Collection("userMessages")
	var bsonFilters bson.D
	for _, filter := range filters {
		bsonFilters = append(bsonFilters, bson.E{
			Key:   filter.Key,
			Value: filter.Value,
		})
	}
	ctx, cancel := repo.MongoDBManager.CreateContext()
	defer cancel()
	cursor, err := collection.Find(ctx, bsonFilters)
	if err != nil {
		log.Error("MongoUserMessageRepo Get error: ", err)
	}
	for cursor.Next(ctx) {
		var result map[string]any
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		UserMessageDTOs = append(UserMessageDTOs, result)
	}
	return UserMessageDTOs
}

func (repo *MongoUserMessageRepo) Create(dto models.UserMessageDTO) (InsertedID string) {
	collection := repo.MongoDBManager.Client.Database("test").Collection("userMessages")
	ctx, cancel := repo.MongoDBManager.CreateContext()
	defer cancel()
	bsonObj, _ := bson.Marshal(dto)
	res, err := collection.InsertOne(ctx, bsonObj)
	if err != nil {
		log.Error("MongoUserMessageRepo create error: ", err)
	}
	InsertedID = res.InsertedID.(primitive.ObjectID).String()
	return InsertedID
}

func NewMongoUserMessageRepo(MongoDBManager *tools.MongoDBManager) *MongoUserMessageRepo {
	return &MongoUserMessageRepo{
		MongoDBManager: MongoDBManager,
	}
}
