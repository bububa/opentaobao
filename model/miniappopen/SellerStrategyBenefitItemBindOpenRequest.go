package miniappopen

import (
	"sync"
)

// SellerStrategyBenefitItemBindOpenRequest 结构体
type SellerStrategyBenefitItemBindOpenRequest struct {
	// C小程序id
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 策略中，实物权益对应的奖池id
	PoolId int64 `json:"pool_id,omitempty" xml:"pool_id,omitempty"`
}

var poolSellerStrategyBenefitItemBindOpenRequest = sync.Pool{
	New: func() any {
		return new(SellerStrategyBenefitItemBindOpenRequest)
	},
}

// GetSellerStrategyBenefitItemBindOpenRequest() 从对象池中获取SellerStrategyBenefitItemBindOpenRequest
func GetSellerStrategyBenefitItemBindOpenRequest() *SellerStrategyBenefitItemBindOpenRequest {
	return poolSellerStrategyBenefitItemBindOpenRequest.Get().(*SellerStrategyBenefitItemBindOpenRequest)
}

// ReleaseSellerStrategyBenefitItemBindOpenRequest 释放SellerStrategyBenefitItemBindOpenRequest
func ReleaseSellerStrategyBenefitItemBindOpenRequest(v *SellerStrategyBenefitItemBindOpenRequest) {
	v.AppId = 0
	v.PoolId = 0
	poolSellerStrategyBenefitItemBindOpenRequest.Put(v)
}
