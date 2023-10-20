package alsc

import (
	"sync"
)

// FeeDetail 结构体
type FeeDetail struct {
	// 计算方式 ：  加-PLUS   加-MINUS
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// MERCHANT_ADJUST  商家调整  PLATFORM_PROM  平台优惠  MERCHANT_PROM  商家优惠  CHARGE  储值  POINT  积分  OTHER  其他
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolFeeDetail = sync.Pool{
	New: func() any {
		return new(FeeDetail)
	},
}

// GetFeeDetail() 从对象池中获取FeeDetail
func GetFeeDetail() *FeeDetail {
	return poolFeeDetail.Get().(*FeeDetail)
}

// ReleaseFeeDetail 释放FeeDetail
func ReleaseFeeDetail(v *FeeDetail) {
	v.Operator = ""
	v.Type = ""
	v.Amount = 0
	poolFeeDetail.Put(v)
}
