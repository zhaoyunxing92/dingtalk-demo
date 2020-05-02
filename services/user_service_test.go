package services

import (
	"fmt"
	"github.com/zhaoyunxing92/dingtalk-demo/config"
	"github.com/zhaoyunxing92/dingtalk-demo/models"
	"github.com/zhaoyunxing92/dingtalk-demo/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

// 添加用户
func TestUserService_Insert(t *testing.T) {

	mongo := config.NewMongoManger("mongodb://localhost:27017", "cherry")
	coll, err := mongo.GetCollection("user")
	if err != nil {
		t.Fatal(err)
	}
	repository := repositories.NewUserRepository(coll)

	service := NewUserService(repository)
	user := models.User{Id: primitive.NewObjectID(), UnionId: "5ead9d8d45628f9c0b3894bd", NickName: "赵云兴", UserName: "zhaoyunxing"}
	id, err := service.Insert(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
}
