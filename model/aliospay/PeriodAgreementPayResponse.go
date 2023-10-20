package aliospay

import (
	"sync"
)

// PeriodAgreementPayResponse 结构体
type PeriodAgreementPayResponse struct {
	// 斑马支付单Id
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
}

var poolPeriodAgreementPayResponse = sync.Pool{
	New: func() any {
		return new(PeriodAgreementPayResponse)
	},
}

// GetPeriodAgreementPayResponse() 从对象池中获取PeriodAgreementPayResponse
func GetPeriodAgreementPayResponse() *PeriodAgreementPayResponse {
	return poolPeriodAgreementPayResponse.Get().(*PeriodAgreementPayResponse)
}

// ReleasePeriodAgreementPayResponse 释放PeriodAgreementPayResponse
func ReleasePeriodAgreementPayResponse(v *PeriodAgreementPayResponse) {
	v.PayOrderId = ""
	poolPeriodAgreementPayResponse.Put(v)
}
