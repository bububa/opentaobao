package ascp

import (
	"sync"
)

// WarehouseCooperationQueryResponse 结构体
type WarehouseCooperationQueryResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回对象
	Data *WarehouseCooperationQueryResultDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWarehouseCooperationQueryResponse = sync.Pool{
	New: func() any {
		return new(WarehouseCooperationQueryResponse)
	},
}

// GetWarehouseCooperationQueryResponse() 从对象池中获取WarehouseCooperationQueryResponse
func GetWarehouseCooperationQueryResponse() *WarehouseCooperationQueryResponse {
	return poolWarehouseCooperationQueryResponse.Get().(*WarehouseCooperationQueryResponse)
}

// ReleaseWarehouseCooperationQueryResponse 释放WarehouseCooperationQueryResponse
func ReleaseWarehouseCooperationQueryResponse(v *WarehouseCooperationQueryResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolWarehouseCooperationQueryResponse.Put(v)
}
