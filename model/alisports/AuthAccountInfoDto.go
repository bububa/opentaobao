package alisports

// AuthAccountInfoDto 结构体
type AuthAccountInfoDto struct {
	// 头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
