package ascp

import (
	"sync"
)

// CombineItemUpsertAsyncResponse 结构体
type CombineItemUpsertAsyncResponse struct {
	// 业务处理结果
	Data []DataItem `json:"data,omitempty" xml:"data>data_item,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=接收失败，1=接收成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCombineItemUpsertAsyncResponse = sync.Pool{
	New: func() any {
		return new(CombineItemUpsertAsyncResponse)
	},
}

// GetCombineItemUpsertAsyncResponse() 从对象池中获取CombineItemUpsertAsyncResponse
func GetCombineItemUpsertAsyncResponse() *CombineItemUpsertAsyncResponse {
	return poolCombineItemUpsertAsyncResponse.Get().(*CombineItemUpsertAsyncResponse)
}

// ReleaseCombineItemUpsertAsyncResponse 释放CombineItemUpsertAsyncResponse
func ReleaseCombineItemUpsertAsyncResponse(v *CombineItemUpsertAsyncResponse) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Result = ""
	v.Success = false
	poolCombineItemUpsertAsyncResponse.Put(v)
}
