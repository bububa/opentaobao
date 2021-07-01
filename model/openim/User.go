package openim

// User 结构体
type User struct {
	// 用户所属appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 是否是淘宝账号
	TaobaoAccount bool `json:"taobao_account,omitempty" xml:"taobao_account,omitempty"`
	// 用户id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 账户appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
}
