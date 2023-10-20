package promotion

import (
	"sync"
)

// CouponTemplateQueryRequest 结构体
type CouponTemplateQueryRequest struct {
	// 模板表ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolCouponTemplateQueryRequest = sync.Pool{
	New: func() any {
		return new(CouponTemplateQueryRequest)
	},
}

// GetCouponTemplateQueryRequest() 从对象池中获取CouponTemplateQueryRequest
func GetCouponTemplateQueryRequest() *CouponTemplateQueryRequest {
	return poolCouponTemplateQueryRequest.Get().(*CouponTemplateQueryRequest)
}

// ReleaseCouponTemplateQueryRequest 释放CouponTemplateQueryRequest
func ReleaseCouponTemplateQueryRequest(v *CouponTemplateQueryRequest) {
	v.Id = 0
	v.SourceId = 0
	v.UserInfo = nil
	poolCouponTemplateQueryRequest.Put(v)
}
