package ascp

import (
	"sync"
)

// WarehouseOperationUpdateResponse 结构体
type WarehouseOperationUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Data *ConsignRuleResultDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWarehouseOperationUpdateResponse = sync.Pool{
	New: func() any {
		return new(WarehouseOperationUpdateResponse)
	},
}

// GetWarehouseOperationUpdateResponse() 从对象池中获取WarehouseOperationUpdateResponse
func GetWarehouseOperationUpdateResponse() *WarehouseOperationUpdateResponse {
	return poolWarehouseOperationUpdateResponse.Get().(*WarehouseOperationUpdateResponse)
}

// ReleaseWarehouseOperationUpdateResponse 释放WarehouseOperationUpdateResponse
func ReleaseWarehouseOperationUpdateResponse(v *WarehouseOperationUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolWarehouseOperationUpdateResponse.Put(v)
}
