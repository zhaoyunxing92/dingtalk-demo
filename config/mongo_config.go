package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type IMongoManager interface {
	GetCollection(name string) (*mongo.Collection, error)
}

type MongoManager struct {
	// 数据库地址
	uri string
	// 数据库名称
	db string
}

func NewMongoManger(uri, db string) IMongoManager {
	return &MongoManager{uri, db}
}

func (mg *MongoManager) GetCollection(name string) (coll *mongo.Collection, err error) {
	client, _ := mongo.NewClient(options.Client().ApplyURI(mg.uri))
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err = client.Connect(ctx)
	coll = client.Database(mg.db).Collection(name)
	return coll, err
}
