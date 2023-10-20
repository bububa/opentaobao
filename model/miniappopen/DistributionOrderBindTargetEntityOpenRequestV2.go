package miniappopen

import (
	"sync"
)

// DistributionOrderBindTargetEntityOpenRequestV2 结构体
type DistributionOrderBindTargetEntityOpenRequestV2 struct {
	// 绑定/解绑的投放计划id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// true:绑定； false: 解绑
	AddBind bool `json:"add_bind,omitempty" xml:"add_bind,omitempty"`
}

var poolDistributionOrderBindTargetEntityOpenRequestV2 = sync.Pool{
	New: func() any {
		return new(DistributionOrderBindTargetEntityOpenRequestV2)
	},
}

// GetDistributionOrderBindTargetEntityOpenRequestV2() 从对象池中获取DistributionOrderBindTargetEntityOpenRequestV2
func GetDistributionOrderBindTargetEntityOpenRequestV2() *DistributionOrderBindTargetEntityOpenRequestV2 {
	return poolDistributionOrderBindTargetEntityOpenRequestV2.Get().(*DistributionOrderBindTargetEntityOpenRequestV2)
}

// ReleaseDistributionOrderBindTargetEntityOpenRequestV2 释放DistributionOrderBindTargetEntityOpenRequestV2
func ReleaseDistributionOrderBindTargetEntityOpenRequestV2(v *DistributionOrderBindTargetEntityOpenRequestV2) {
	v.OrderId = 0
	v.AddBind = false
	poolDistributionOrderBindTargetEntityOpenRequestV2.Put(v)
}
