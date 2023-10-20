package promotion

import (
	"sync"
)

// CouponTemplateTerminateRequest 结构体
type CouponTemplateTerminateRequest struct {
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolCouponTemplateTerminateRequest = sync.Pool{
	New: func() any {
		return new(CouponTemplateTerminateRequest)
	},
}

// GetCouponTemplateTerminateRequest() 从对象池中获取CouponTemplateTerminateRequest
func GetCouponTemplateTerminateRequest() *CouponTemplateTerminateRequest {
	return poolCouponTemplateTerminateRequest.Get().(*CouponTemplateTerminateRequest)
}

// ReleaseCouponTemplateTerminateRequest 释放CouponTemplateTerminateRequest
func ReleaseCouponTemplateTerminateRequest(v *CouponTemplateTerminateRequest) {
	v.SourceId = 0
	v.UserInfo = nil
	poolCouponTemplateTerminateRequest.Put(v)
}
