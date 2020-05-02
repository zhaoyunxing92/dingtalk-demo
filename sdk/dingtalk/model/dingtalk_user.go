package model

/**
钉钉用户信息
*/
type UserInfo struct {
	/**
	用户在钉钉上面的昵称
	*/
	Nick string `json:"nick"`
	/**
	用户在当前开放应用内的唯一标识
	*/
	Openid string `json:"openid"`
	/**
	用户在当前开放应用所属企业内的唯一标识
	*/
	Unionid string `json:"unionid"`
}

/**
钉钉返回结果  "access_token": "fw8ef8we8f76e6f7s8df8s"
*/
type Response struct {
	Code        int      `json:"errcode"`
	Msg         string   `json:"errmsg"`
	UserInfo    UserInfo `json:"user_info"`
}