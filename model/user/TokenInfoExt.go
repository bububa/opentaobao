package user

import (
	"sync"
)

// TokenInfoExt 结构体
type TokenInfoExt struct {
	// open account当前token info中open account id对应的open account信息
	OpenAccount *OpenAccount `json:"open_account,omitempty" xml:"open_account,omitempty"`
	// oauthOtherInfo
	OauthOtherInfo *OAuthOtherInfo `json:"oauth_other_info,omitempty" xml:"oauth_other_info,omitempty"`
}

var poolTokenInfoExt = sync.Pool{
	New: func() any {
		return new(TokenInfoExt)
	},
}

// GetTokenInfoExt() 从对象池中获取TokenInfoExt
func GetTokenInfoExt() *TokenInfoExt {
	return poolTokenInfoExt.Get().(*TokenInfoExt)
}

// ReleaseTokenInfoExt 释放TokenInfoExt
func ReleaseTokenInfoExt(v *TokenInfoExt) {
	v.OpenAccount = nil
	v.OauthOtherInfo = nil
	poolTokenInfoExt.Put(v)
}
