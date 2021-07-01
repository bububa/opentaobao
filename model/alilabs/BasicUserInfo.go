package alilabs

// BasicUserInfo 结构体
type BasicUserInfo struct {
	// 头像 url
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
}
