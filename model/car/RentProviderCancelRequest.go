package car

import (
	"sync"
)

// RentProviderCancelRequest 结构体
type RentProviderCancelRequest struct {
	// 取消拒绝原因
	CancelRejectReason string `json:"cancel_reject_reason,omitempty" xml:"cancel_reject_reason,omitempty"`
	// 服务商ID
	ProviderId int64 `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 取消拒绝类型
	CancelRejectCode int64 `json:"cancel_reject_code,omitempty" xml:"cancel_reject_code,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 是否确认可以取消
	CancelConfirm bool `json:"cancel_confirm,omitempty" xml:"cancel_confirm,omitempty"`
}

var poolRentProviderCancelRequest = sync.Pool{
	New: func() any {
		return new(RentProviderCancelRequest)
	},
}

// GetRentProviderCancelRequest() 从对象池中获取RentProviderCancelRequest
func GetRentProviderCancelRequest() *RentProviderCancelRequest {
	return poolRentProviderCancelRequest.Get().(*RentProviderCancelRequest)
}

// ReleaseRentProviderCancelRequest 释放RentProviderCancelRequest
func ReleaseRentProviderCancelRequest(v *RentProviderCancelRequest) {
	v.CancelRejectReason = ""
	v.ProviderId = 0
	v.CancelRejectCode = 0
	v.OrderId = 0
	v.CancelConfirm = false
	poolRentProviderCancelRequest.Put(v)
}
