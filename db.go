package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 持久化接口,存对数据进行持久化
type Persistence interface {
	Save(documents []interface{}) (mongo.InsertManyResult, error)
	SaveOne(document interface{}) (mongo.InsertOneResult, error)
}

// 创建mongo客户端
func CreateMonoPersistence() Persistence {
	mongoClient := new(mongoPersistence)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	mongoClient.collection = client.Database("douban").Collection("tv")
	return mongoClient
}

type mongoPersistence struct {
	collection *mongo.Collection
}

func (this *mongoPersistence) Save(documents []interface{}) (mongo.InsertManyResult, error) {
	res, err := this.collection.InsertMany(context.TODO(), documents)
	return *res, err
}

func (this *mongoPersistence) SaveOne(document interface{}) (mongo.InsertOneResult, error) {
	res, err := this.collection.InsertOne(context.TODO(), document)
	return *res, err
}