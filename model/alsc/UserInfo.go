package alsc

// UserInfo 结构体
type UserInfo struct {
	// 账号类型，0表示淘宝账号，25表示饿了么账号
	Site string `json:"site,omitempty" xml:"site,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
