package omniorder

import (
	"sync"
)

// SdtStatusResponse 结构体
type SdtStatusResponse struct {
	// 取消原因
	ReasonDesc string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
	// 卖家ID，通sellerID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// packageId
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 状态 0 取号，1 已发货 -1 商家取消 -2 运力端取消
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSdtStatusResponse = sync.Pool{
	New: func() any {
		return new(SdtStatusResponse)
	},
}

// GetSdtStatusResponse() 从对象池中获取SdtStatusResponse
func GetSdtStatusResponse() *SdtStatusResponse {
	return poolSdtStatusResponse.Get().(*SdtStatusResponse)
}

// ReleaseSdtStatusResponse 释放SdtStatusResponse
func ReleaseSdtStatusResponse(v *SdtStatusResponse) {
	v.ReasonDesc = ""
	v.UserId = 0
	v.OrderId = 0
	v.Status = 0
	poolSdtStatusResponse.Put(v)
}
