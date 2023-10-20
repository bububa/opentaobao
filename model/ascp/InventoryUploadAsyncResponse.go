package ascp

import (
	"sync"
)

// InventoryUploadAsyncResponse 结构体
type InventoryUploadAsyncResponse struct {
	// 业务处理结果
	Data []DataDetail `json:"data,omitempty" xml:"data>data_detail,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=接收失败，1=接收成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolInventoryUploadAsyncResponse = sync.Pool{
	New: func() any {
		return new(InventoryUploadAsyncResponse)
	},
}

// GetInventoryUploadAsyncResponse() 从对象池中获取InventoryUploadAsyncResponse
func GetInventoryUploadAsyncResponse() *InventoryUploadAsyncResponse {
	return poolInventoryUploadAsyncResponse.Get().(*InventoryUploadAsyncResponse)
}

// ReleaseInventoryUploadAsyncResponse 释放InventoryUploadAsyncResponse
func ReleaseInventoryUploadAsyncResponse(v *InventoryUploadAsyncResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolInventoryUploadAsyncResponse.Put(v)
}
