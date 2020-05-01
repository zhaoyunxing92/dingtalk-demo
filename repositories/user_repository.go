package repositories

import (
	"context"
	"github.com/zhaoyunxing90/dingtalk-demo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"gopkg.in/ini.v1"
	"time"
)

type UserRepository interface {
	/**
	连接数据库
	*/
	Conn() error
	/**
	添加用户
	*/
	Insert(user *models.User) (string, error)
	/**
	更新用户信息
	*/
	Update(user *models.User) (result *mongo.UpdateResult, err error)
	/**
	  根据id获取用户
	*/
	SelectById(string) (*models.User, error)

	/**
	根据openid获取用户
	*/
	//SelectByOpenId(string) (*models.User, error)
}

type UserManager struct {
	/**
	集合名称
	*/
	table string
	/**
	配置地址
	*/
	cfg string

	/**
	集合对象
	*/
	coll *mongo.Collection
}

func NewUserRepository(cfg string) UserRepository {

	return &UserManager{"user", cfg, nil}
}

// 连接
func (manger *UserManager) Conn() (err error) {

	if manger.coll == nil {
		cfg, err := ini.Load(manger.cfg)
		if err != nil {
			return err
		}
		uri := cfg.Section("mongo").Key("uri").String()
		database := cfg.Section("mongo").Key("database").String()
		client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			return err
		}
		manger.coll = client.Database(database).Collection(manger.table)
	}
	return
}

// 添加
func (manger *UserManager) Insert(user *models.User) (id string, err error) {
	//if err = manger.Conn(); err == nil {
	//	return
	//}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := manger.coll.InsertOne(ctx, user)
	if err != nil {
		return
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

// 添加
func (manger *UserManager) SelectById(id string) (user *models.User, err error) {
	//if err = manger.Conn(); err == nil {
	//	return user, nil
	//}
	oldId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": bsonx.ObjectID(oldId)}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = manger.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return
	}
	return user, err
}

/**
https://godoc.org/go.mongodb.org/mongo-driver/mongo#Collection.UpdateOne
*/
// 根据id更新
func (manger *UserManager) Update(user *models.User) (result *mongo.UpdateResult, err error) {
	// 存在就修改否则就插入
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": user.GetId()}
	update := bson.D{{"$set", user}}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return manger.coll.UpdateOne(ctx, filter, update, opts)
}
