package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk-demo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestUserManager_Insert(t *testing.T) {

	repository := NewUserRepositoryCfg("D:/github/go/dingtalk-demo/cfg.ini")
	err := repository.Conn()
	if err != nil {
		t.Fatal(err)
	}
	user := models.User{Id: primitive.NewObjectID(), ModifyTime: time.Now().Format("2006-01-02 15:04:05"), CreateTime: time.Now().Format("2006-01-02 15:04:05"), NickName: "赵云兴", UserName: "zhaoyunxing"}
	id, err := repository.Insert(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)

}

func TestUserManager_SelectById(t *testing.T) {
	repository := NewUserRepositoryCfg("D:/github/go/dingtalk-demo/cfg.ini")
	err := repository.Conn()
	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.SelectById("5eac6317cbc997dda73c286f")
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(marshal))
}

func TestUserManager_Update(t *testing.T) {
	repository := NewUserRepositoryCfg("D:/github/go/dingtalk-demo/cfg.ini")
	err := repository.Conn()
	if err != nil {
		t.Fatal(err)
	}
	user := models.User{NickName: "赵子龙", UserName: "zhaoyunxing"}
	user.SetId("5eac63b7ba4dca9c11b5ec89")

	res, err := repository.Update(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
