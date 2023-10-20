package wdk

import (
	"sync"
)

// PlatformDeduction 结构体
type PlatformDeduction struct {
	// 技术服务费
	TechnicalServiceFee int64 `json:"technical_service_fee,omitempty" xml:"technical_service_fee,omitempty"`
	// 支付服务费
	PayServiceFee int64 `json:"pay_service_fee,omitempty" xml:"pay_service_fee,omitempty"`
	// 基础物流费
	BaseLogisticsFee int64 `json:"base_logistics_fee,omitempty" xml:"base_logistics_fee,omitempty"`
	// 代运营服务费
	ThirdpartnarFee int64 `json:"thirdpartnar_fee,omitempty" xml:"thirdpartnar_fee,omitempty"`
}

var poolPlatformDeduction = sync.Pool{
	New: func() any {
		return new(PlatformDeduction)
	},
}

// GetPlatformDeduction() 从对象池中获取PlatformDeduction
func GetPlatformDeduction() *PlatformDeduction {
	return poolPlatformDeduction.Get().(*PlatformDeduction)
}

// ReleasePlatformDeduction 释放PlatformDeduction
func ReleasePlatformDeduction(v *PlatformDeduction) {
	v.TechnicalServiceFee = 0
	v.PayServiceFee = 0
	v.BaseLogisticsFee = 0
	v.ThirdpartnarFee = 0
	poolPlatformDeduction.Put(v)
}
