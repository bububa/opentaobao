package ascp

import (
	"sync"
)

// ItemDeleteAsyncResponse 结构体
type ItemDeleteAsyncResponse struct {
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

var poolItemDeleteAsyncResponse = sync.Pool{
	New: func() any {
		return new(ItemDeleteAsyncResponse)
	},
}

// GetItemDeleteAsyncResponse() 从对象池中获取ItemDeleteAsyncResponse
func GetItemDeleteAsyncResponse() *ItemDeleteAsyncResponse {
	return poolItemDeleteAsyncResponse.Get().(*ItemDeleteAsyncResponse)
}

// ReleaseItemDeleteAsyncResponse 释放ItemDeleteAsyncResponse
func ReleaseItemDeleteAsyncResponse(v *ItemDeleteAsyncResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolItemDeleteAsyncResponse.Put(v)
}
