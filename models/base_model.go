package models

type BaseModel struct {
	Id string `json:"id"`
	/**
	状态
	*/
	State string `json:"state"`
	/**
	创建时间
	*/
	CreateTime string `json:"createTime"`
	/**
	修改时间
	*/
	ModifyTime string `json:"modifyTime"`
}
