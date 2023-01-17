package alsc

// UserInfo 结构体
type UserInfo struct {
	// 用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 账号类型，0表示淘宝账号，25表示饿了么账号
	Site int64 `json:"site,omitempty" xml:"site,omitempty"`
}
