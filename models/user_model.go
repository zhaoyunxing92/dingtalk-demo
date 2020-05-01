package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
)

type User struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	/**
	状态
	*/
	State string `json:"state"`
	/**
	用户名
	*/
	UserName string `json:"userName"`
	/**
	昵称 长度为1~64个字符
	*/
	NickName string `json:"nickName"`
	/**
	用户在当前开放应用内的唯一标识
	*/
	OpenId string `json:"openId"`
	/**
	用户在当前开放应用所属企业内的唯一标识
	*/
	UnionId string `json:"unionId"`
	/**
	  用户头像
	*/
	Avatar string `json:"avatar"`
	/**
	备注
	*/
	Remark string `json:"avatar"`
	/**
	邮箱
	*/
	Email string `json:"email"`
	/**
	orgEmail
	*/
	OrgEmail string `json:"orgEmail"`
	/**
	职位信息。长度为0~64个字符
	*/
	Position string `json:"position"`
	/**
	创建时间
	*/
	CreateTime string `json:"createTime"`
	/**
	修改时间
	*/
	ModifyTime string `json:"modifyTime"`
}

/**
设置 objectId
*/
func (user *User) SetId(id string) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	user.Id = objectId
	return
}

/**
获取id
*/
func (user *User) GetId() bsonx.Val {

	return bsonx.ObjectID(user.Id)
}
