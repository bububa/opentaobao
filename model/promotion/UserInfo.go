package promotion

// UserInfo 结构体
type UserInfo struct {
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}
