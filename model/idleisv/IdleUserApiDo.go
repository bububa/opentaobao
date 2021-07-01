package idleisv

// IdleUserApiDo 结构体
type IdleUserApiDo struct {
	// 淘宝用户Nick
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户身份（GENERAL: 普通用户，PRO_PLAYER: 个人经营者）
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
}
