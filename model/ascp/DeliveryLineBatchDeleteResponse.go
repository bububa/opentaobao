package ascp

import (
	"sync"
)

// DeliveryLineBatchDeleteResponse 结构体
type DeliveryLineBatchDeleteResponse struct {
	// 线路删除失败信息
	Data []SignRuleDeleteResultDetail `json:"data,omitempty" xml:"data>sign_rule_delete_result_detail,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryLineBatchDeleteResponse = sync.Pool{
	New: func() any {
		return new(DeliveryLineBatchDeleteResponse)
	},
}

// GetDeliveryLineBatchDeleteResponse() 从对象池中获取DeliveryLineBatchDeleteResponse
func GetDeliveryLineBatchDeleteResponse() *DeliveryLineBatchDeleteResponse {
	return poolDeliveryLineBatchDeleteResponse.Get().(*DeliveryLineBatchDeleteResponse)
}

// ReleaseDeliveryLineBatchDeleteResponse 释放DeliveryLineBatchDeleteResponse
func ReleaseDeliveryLineBatchDeleteResponse(v *DeliveryLineBatchDeleteResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolDeliveryLineBatchDeleteResponse.Put(v)
}
