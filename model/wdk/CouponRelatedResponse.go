package wdk

import (
	"sync"
)

// CouponRelatedResponse 结构体
type CouponRelatedResponse struct {
	// 活动id
	UmpActivityId int64 `json:"ump_activity_id,omitempty" xml:"ump_activity_id,omitempty"`
	// 券模版id
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
}

var poolCouponRelatedResponse = sync.Pool{
	New: func() any {
		return new(CouponRelatedResponse)
	},
}

// GetCouponRelatedResponse() 从对象池中获取CouponRelatedResponse
func GetCouponRelatedResponse() *CouponRelatedResponse {
	return poolCouponRelatedResponse.Get().(*CouponRelatedResponse)
}

// ReleaseCouponRelatedResponse 释放CouponRelatedResponse
func ReleaseCouponRelatedResponse(v *CouponRelatedResponse) {
	v.UmpActivityId = 0
	v.SourceId = 0
	poolCouponRelatedResponse.Put(v)
}
