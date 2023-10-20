package ascp

import (
	"sync"
)

// DeliveryDecisionQueryResponse 结构体
type DeliveryDecisionQueryResponse struct {
	// 详细信息
	Data []DataDetail `json:"data,omitempty" xml:"data>data_detail,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryDecisionQueryResponse = sync.Pool{
	New: func() any {
		return new(DeliveryDecisionQueryResponse)
	},
}

// GetDeliveryDecisionQueryResponse() 从对象池中获取DeliveryDecisionQueryResponse
func GetDeliveryDecisionQueryResponse() *DeliveryDecisionQueryResponse {
	return poolDeliveryDecisionQueryResponse.Get().(*DeliveryDecisionQueryResponse)
}

// ReleaseDeliveryDecisionQueryResponse 释放DeliveryDecisionQueryResponse
func ReleaseDeliveryDecisionQueryResponse(v *DeliveryDecisionQueryResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolDeliveryDecisionQueryResponse.Put(v)
}
