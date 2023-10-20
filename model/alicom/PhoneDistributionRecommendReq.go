package alicom

import (
	"sync"
)

// PhoneDistributionRecommendReq 结构体
type PhoneDistributionRecommendReq struct {
	// 用户手机号
	Account string `json:"account,omitempty" xml:"account,omitempty"`
}

var poolPhoneDistributionRecommendReq = sync.Pool{
	New: func() any {
		return new(PhoneDistributionRecommendReq)
	},
}

// GetPhoneDistributionRecommendReq() 从对象池中获取PhoneDistributionRecommendReq
func GetPhoneDistributionRecommendReq() *PhoneDistributionRecommendReq {
	return poolPhoneDistributionRecommendReq.Get().(*PhoneDistributionRecommendReq)
}

// ReleasePhoneDistributionRecommendReq 释放PhoneDistributionRecommendReq
func ReleasePhoneDistributionRecommendReq(v *PhoneDistributionRecommendReq) {
	v.Account = ""
	poolPhoneDistributionRecommendReq.Put(v)
}
