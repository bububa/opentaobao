package tmallcar

import (
	"sync"
)

// AuthCheckReq 结构体
type AuthCheckReq struct {
	// 鉴权token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolAuthCheckReq = sync.Pool{
	New: func() any {
		return new(AuthCheckReq)
	},
}

// GetAuthCheckReq() 从对象池中获取AuthCheckReq
func GetAuthCheckReq() *AuthCheckReq {
	return poolAuthCheckReq.Get().(*AuthCheckReq)
}

// ReleaseAuthCheckReq 释放AuthCheckReq
func ReleaseAuthCheckReq(v *AuthCheckReq) {
	v.Token = ""
	v.OrderId = 0
	poolAuthCheckReq.Put(v)
}
