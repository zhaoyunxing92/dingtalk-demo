package services

import (
	"github.com/zhaoyunxing92/dingtalk-demo/models"
	"github.com/zhaoyunxing92/dingtalk-demo/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IUserService interface {
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
	SelectByOpenId(string) (*models.User, error)

	/**
	根据unionId获取用户
	*/
	SelectByUnionId(string) (*models.User, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

// 初始化函数
func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

/**
添加用户
*/
func (service *UserService) Insert(user *models.User) (string, error) {
	//给时间默认值
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return service.repository.Insert(user)
}

/**
更新用户
*/
func (service *UserService) Update(user *models.User) (result *mongo.UpdateResult, err error) {
	//设置时间
	user.ModifyTime = time.Now().Format("2006-01-02 15:04:05")
	return service.repository.Update(user)
}

/**
根据id获取用户
*/
func (service *UserService) SelectById(id string) (*models.User, error) {
	//user, err := service.repository.SelectById(id)
	//if err != nil {
	//	return "", err
	//}
	//marshal, err := json.Marshal(user)
	//if err != nil {
	//	return "", err
	//}
	//str = string(marshal)
	return service.repository.SelectById(id)
}

/**
根据openid获取用户
*/
func (service *UserService) SelectByOpenId(openid string) (*models.User, error) {
	return service.repository.SelectByOpenId(openid)
}

/**
根据unionId获取用户
*/
func (service *UserService) SelectByUnionId(unionid string) (*models.User, error) {
	return service.repository.SelectByUnionId(unionid)
}
