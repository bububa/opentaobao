package ascp

import (
	"sync"
)

// WarehouseCooperationUpdateResponse 结构体
type WarehouseCooperationUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回内容
	Data *WarehouseCooperationResultDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWarehouseCooperationUpdateResponse = sync.Pool{
	New: func() any {
		return new(WarehouseCooperationUpdateResponse)
	},
}

// GetWarehouseCooperationUpdateResponse() 从对象池中获取WarehouseCooperationUpdateResponse
func GetWarehouseCooperationUpdateResponse() *WarehouseCooperationUpdateResponse {
	return poolWarehouseCooperationUpdateResponse.Get().(*WarehouseCooperationUpdateResponse)
}

// ReleaseWarehouseCooperationUpdateResponse 释放WarehouseCooperationUpdateResponse
func ReleaseWarehouseCooperationUpdateResponse(v *WarehouseCooperationUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolWarehouseCooperationUpdateResponse.Put(v)
}
