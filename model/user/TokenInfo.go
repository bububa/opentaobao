package user

import (
	"sync"
)

// TokenInfo 结构体
type TokenInfo struct {
	// isv自己账号的唯一id
	IsvAccountId string `json:"isv_account_id,omitempty" xml:"isv_account_id,omitempty"`
	// 用于防重放的唯一id
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// token info扩展信息
	Ext *TokenInfoExt `json:"ext,omitempty" xml:"ext,omitempty"`
	// ISV APP的登录态时长
	LoginStateExpireIn int64 `json:"login_state_expire_in,omitempty" xml:"login_state_expire_in,omitempty"`
	// open account id
	OpenAccountId int64 `json:"open_account_id,omitempty" xml:"open_account_id,omitempty"`
	// 时间戳
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

var poolTokenInfo = sync.Pool{
	New: func() any {
		return new(TokenInfo)
	},
}

// GetTokenInfo() 从对象池中获取TokenInfo
func GetTokenInfo() *TokenInfo {
	return poolTokenInfo.Get().(*TokenInfo)
}

// ReleaseTokenInfo 释放TokenInfo
func ReleaseTokenInfo(v *TokenInfo) {
	v.IsvAccountId = ""
	v.Uuid = ""
	v.Ext = nil
	v.LoginStateExpireIn = 0
	v.OpenAccountId = 0
	v.Timestamp = 0
	poolTokenInfo.Put(v)
}
