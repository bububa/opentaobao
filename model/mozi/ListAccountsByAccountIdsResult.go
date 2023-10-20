package mozi

import (
	"sync"
)

// ListAccountsByAccountIdsResult 结构体
type ListAccountsByAccountIdsResult struct {
	// 返回的数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求的序列化
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的状态消息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 返回的附件信息
	ResponseMetaData string `json:"response_meta_data,omitempty" xml:"response_meta_data,omitempty"`
	// 返回码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolListAccountsByAccountIdsResult = sync.Pool{
	New: func() any {
		return new(ListAccountsByAccountIdsResult)
	},
}

// GetListAccountsByAccountIdsResult() 从对象池中获取ListAccountsByAccountIdsResult
func GetListAccountsByAccountIdsResult() *ListAccountsByAccountIdsResult {
	return poolListAccountsByAccountIdsResult.Get().(*ListAccountsByAccountIdsResult)
}

// ReleaseListAccountsByAccountIdsResult 释放ListAccountsByAccountIdsResult
func ReleaseListAccountsByAccountIdsResult(v *ListAccountsByAccountIdsResult) {
	v.Data = ""
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseMetaData = ""
	v.ResponseCode = ""
	v.Success = false
	poolListAccountsByAccountIdsResult.Put(v)
}
