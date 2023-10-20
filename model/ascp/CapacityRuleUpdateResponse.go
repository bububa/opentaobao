package ascp

import (
	"sync"
)

// CapacityRuleUpdateResponse 结构体
type CapacityRuleUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCapacityRuleUpdateResponse = sync.Pool{
	New: func() any {
		return new(CapacityRuleUpdateResponse)
	},
}

// GetCapacityRuleUpdateResponse() 从对象池中获取CapacityRuleUpdateResponse
func GetCapacityRuleUpdateResponse() *CapacityRuleUpdateResponse {
	return poolCapacityRuleUpdateResponse.Get().(*CapacityRuleUpdateResponse)
}

// ReleaseCapacityRuleUpdateResponse 释放CapacityRuleUpdateResponse
func ReleaseCapacityRuleUpdateResponse(v *CapacityRuleUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolCapacityRuleUpdateResponse.Put(v)
}
