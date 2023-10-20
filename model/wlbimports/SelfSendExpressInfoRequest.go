package wlbimports

import (
	"sync"
)

// SelfSendExpressInfoRequest 结构体
type SelfSendExpressInfoRequest struct {
	// 快递运单号（自寄快递模式必填）
	ExpressNumber string `json:"express_number,omitempty" xml:"express_number,omitempty"`
	// 快递公司（自寄快递模式必填）
	ExpressCompany string `json:"express_company,omitempty" xml:"express_company,omitempty"`
	// 预计送达时间（自寄快递模式必填）
	ExpectedArriveTime int64 `json:"expected_arrive_time,omitempty" xml:"expected_arrive_time,omitempty"`
	// 预计揽收时间（自寄快递模式必填）
	ExpectedPickupTime int64 `json:"expected_pickup_time,omitempty" xml:"expected_pickup_time,omitempty"`
}

var poolSelfSendExpressInfoRequest = sync.Pool{
	New: func() any {
		return new(SelfSendExpressInfoRequest)
	},
}

// GetSelfSendExpressInfoRequest() 从对象池中获取SelfSendExpressInfoRequest
func GetSelfSendExpressInfoRequest() *SelfSendExpressInfoRequest {
	return poolSelfSendExpressInfoRequest.Get().(*SelfSendExpressInfoRequest)
}

// ReleaseSelfSendExpressInfoRequest 释放SelfSendExpressInfoRequest
func ReleaseSelfSendExpressInfoRequest(v *SelfSendExpressInfoRequest) {
	v.ExpressNumber = ""
	v.ExpressCompany = ""
	v.ExpectedArriveTime = 0
	v.ExpectedPickupTime = 0
	poolSelfSendExpressInfoRequest.Put(v)
}
