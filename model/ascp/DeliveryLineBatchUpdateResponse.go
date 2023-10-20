package ascp

import (
	"sync"
)

// DeliveryLineBatchUpdateResponse 结构体
type DeliveryLineBatchUpdateResponse struct {
	// 处理失败的到货线路规则（组）
	Data []SignRuleResultDetail `json:"data,omitempty" xml:"data>sign_rule_result_detail,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryLineBatchUpdateResponse = sync.Pool{
	New: func() any {
		return new(DeliveryLineBatchUpdateResponse)
	},
}

// GetDeliveryLineBatchUpdateResponse() 从对象池中获取DeliveryLineBatchUpdateResponse
func GetDeliveryLineBatchUpdateResponse() *DeliveryLineBatchUpdateResponse {
	return poolDeliveryLineBatchUpdateResponse.Get().(*DeliveryLineBatchUpdateResponse)
}

// ReleaseDeliveryLineBatchUpdateResponse 释放DeliveryLineBatchUpdateResponse
func ReleaseDeliveryLineBatchUpdateResponse(v *DeliveryLineBatchUpdateResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolDeliveryLineBatchUpdateResponse.Put(v)
}
