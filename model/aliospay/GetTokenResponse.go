package aliospay

import (
	"sync"
)

// GetTokenResponse 结构体
type GetTokenResponse struct {
	// 支付token
	PayToken string `json:"pay_token,omitempty" xml:"pay_token,omitempty"`
}

var poolGetTokenResponse = sync.Pool{
	New: func() any {
		return new(GetTokenResponse)
	},
}

// GetGetTokenResponse() 从对象池中获取GetTokenResponse
func GetGetTokenResponse() *GetTokenResponse {
	return poolGetTokenResponse.Get().(*GetTokenResponse)
}

// ReleaseGetTokenResponse 释放GetTokenResponse
func ReleaseGetTokenResponse(v *GetTokenResponse) {
	v.PayToken = ""
	poolGetTokenResponse.Put(v)
}
