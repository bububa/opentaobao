package miniappopen

import (
	"sync"
)

// DistributionOrderQueryByIdOpenRequest 结构体
type DistributionOrderQueryByIdOpenRequest struct {
	// 投放计划的id列表
	OrderIdList []int64 `json:"order_id_list,omitempty" xml:"order_id_list>int64,omitempty"`
}

var poolDistributionOrderQueryByIdOpenRequest = sync.Pool{
	New: func() any {
		return new(DistributionOrderQueryByIdOpenRequest)
	},
}

// GetDistributionOrderQueryByIdOpenRequest() 从对象池中获取DistributionOrderQueryByIdOpenRequest
func GetDistributionOrderQueryByIdOpenRequest() *DistributionOrderQueryByIdOpenRequest {
	return poolDistributionOrderQueryByIdOpenRequest.Get().(*DistributionOrderQueryByIdOpenRequest)
}

// ReleaseDistributionOrderQueryByIdOpenRequest 释放DistributionOrderQueryByIdOpenRequest
func ReleaseDistributionOrderQueryByIdOpenRequest(v *DistributionOrderQueryByIdOpenRequest) {
	v.OrderIdList = v.OrderIdList[:0]
	poolDistributionOrderQueryByIdOpenRequest.Put(v)
}
