package logistic

import (
	"sync"
)

// Tms2MscPayQueryResponse 结构体
type Tms2MscPayQueryResponse struct {
	// 消费者支付状态
	PayFlag string `json:"pay_flag,omitempty" xml:"pay_flag,omitempty"`
	// 消费者支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
}

var poolTms2MscPayQueryResponse = sync.Pool{
	New: func() any {
		return new(Tms2MscPayQueryResponse)
	},
}

// GetTms2MscPayQueryResponse() 从对象池中获取Tms2MscPayQueryResponse
func GetTms2MscPayQueryResponse() *Tms2MscPayQueryResponse {
	return poolTms2MscPayQueryResponse.Get().(*Tms2MscPayQueryResponse)
}

// ReleaseTms2MscPayQueryResponse 释放Tms2MscPayQueryResponse
func ReleaseTms2MscPayQueryResponse(v *Tms2MscPayQueryResponse) {
	v.PayFlag = ""
	v.PayTime = ""
	poolTms2MscPayQueryResponse.Put(v)
}
