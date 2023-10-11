package idle

// IdleGooseFishUserInfoVo 结构体
type IdleGooseFishUserInfoVo struct {
	// 闲鱼头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 用户闲鱼昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 0男生 1女生
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}
