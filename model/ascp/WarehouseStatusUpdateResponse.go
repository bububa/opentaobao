package ascp

import (
	"sync"
)

// WarehouseStatusUpdateResponse 结构体
type WarehouseStatusUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWarehouseStatusUpdateResponse = sync.Pool{
	New: func() any {
		return new(WarehouseStatusUpdateResponse)
	},
}

// GetWarehouseStatusUpdateResponse() 从对象池中获取WarehouseStatusUpdateResponse
func GetWarehouseStatusUpdateResponse() *WarehouseStatusUpdateResponse {
	return poolWarehouseStatusUpdateResponse.Get().(*WarehouseStatusUpdateResponse)
}

// ReleaseWarehouseStatusUpdateResponse 释放WarehouseStatusUpdateResponse
func ReleaseWarehouseStatusUpdateResponse(v *WarehouseStatusUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolWarehouseStatusUpdateResponse.Put(v)
}
