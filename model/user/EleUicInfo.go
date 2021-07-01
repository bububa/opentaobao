package user

// EleUicInfo 结构体
type EleUicInfo struct {
	// 用户头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
