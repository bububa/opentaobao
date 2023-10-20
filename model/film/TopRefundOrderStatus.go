package film

import (
	"sync"
)

// TopRefundOrderStatus 结构体
type TopRefundOrderStatus struct {
	// 退款中，其他状态可详见接口文档
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// tbOrderId
	TbOrderId string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTopRefundOrderStatus = sync.Pool{
	New: func() any {
		return new(TopRefundOrderStatus)
	},
}

// GetTopRefundOrderStatus() 从对象池中获取TopRefundOrderStatus
func GetTopRefundOrderStatus() *TopRefundOrderStatus {
	return poolTopRefundOrderStatus.Get().(*TopRefundOrderStatus)
}

// ReleaseTopRefundOrderStatus 释放TopRefundOrderStatus
func ReleaseTopRefundOrderStatus(v *TopRefundOrderStatus) {
	v.Status = ""
	v.TbOrderId = ""
	v.Message = ""
	poolTopRefundOrderStatus.Put(v)
}
