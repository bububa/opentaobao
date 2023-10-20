package mos

import (
	"sync"
)

// CallDispatcherResponse 结构体
type CallDispatcherResponse struct {
	// respList
	RespList []CallDispatcherRespDo `json:"resp_list,omitempty" xml:"resp_list>call_dispatcher_resp_do,omitempty"`
	// parentNo
	ParentNo string `json:"parent_no,omitempty" xml:"parent_no,omitempty"`
}

var poolCallDispatcherResponse = sync.Pool{
	New: func() any {
		return new(CallDispatcherResponse)
	},
}

// GetCallDispatcherResponse() 从对象池中获取CallDispatcherResponse
func GetCallDispatcherResponse() *CallDispatcherResponse {
	return poolCallDispatcherResponse.Get().(*CallDispatcherResponse)
}

// ReleaseCallDispatcherResponse 释放CallDispatcherResponse
func ReleaseCallDispatcherResponse(v *CallDispatcherResponse) {
	v.RespList = v.RespList[:0]
	v.ParentNo = ""
	poolCallDispatcherResponse.Put(v)
}
