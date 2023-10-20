package alitripmerchant

import (
	"sync"
)

// WechatCodeResponse 结构体
type WechatCodeResponse struct {
	// 微信union_Id
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 微信用户openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolWechatCodeResponse = sync.Pool{
	New: func() any {
		return new(WechatCodeResponse)
	},
}

// GetWechatCodeResponse() 从对象池中获取WechatCodeResponse
func GetWechatCodeResponse() *WechatCodeResponse {
	return poolWechatCodeResponse.Get().(*WechatCodeResponse)
}

// ReleaseWechatCodeResponse 释放WechatCodeResponse
func ReleaseWechatCodeResponse(v *WechatCodeResponse) {
	v.UnionId = ""
	v.OpenId = ""
	poolWechatCodeResponse.Put(v)
}
