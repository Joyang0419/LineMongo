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

func (repo *MongoUserMessageRepo) Get() (UserMessageDTOs []models.UserMessageDTO) {
	collection := repo.MongoDBManager.Client.Database("test").Collection("userMessages")
	filter := bson.D{{"replytoken", "hello"}}
	ctx, cancel := repo.MongoDBManager.CreateContext()
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error("MongoUserMessageRepo Get error: ", err)
	}
	for cursor.Next(ctx) {
		var result models.UserMessageDTO
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
