package crm

import (
	"sync"
)

// TaobaoCrmMemberJoinurlGetResult 结构体
type TaobaoCrmMemberJoinurlGetResult struct {
	// 返回入会地址URL，需自行判断协议头，如返回结果为//建议使用https://
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	//
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// exceptionCode
	ExceptionCode string `json:"exception_code,omitempty" xml:"exception_code,omitempty"`
	// total
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoCrmMemberJoinurlGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoCrmMemberJoinurlGetResult)
	},
}

// GetTaobaoCrmMemberJoinurlGetResult() 从对象池中获取TaobaoCrmMemberJoinurlGetResult
func GetTaobaoCrmMemberJoinurlGetResult() *TaobaoCrmMemberJoinurlGetResult {
	return poolTaobaoCrmMemberJoinurlGetResult.Get().(*TaobaoCrmMemberJoinurlGetResult)
}

// ReleaseTaobaoCrmMemberJoinurlGetResult 释放TaobaoCrmMemberJoinurlGetResult
func ReleaseTaobaoCrmMemberJoinurlGetResult(v *TaobaoCrmMemberJoinurlGetResult) {
	v.Result = ""
	v.ErrorMsg = ""
	v.ExceptionCode = ""
	v.Total = 0
	v.ErrorCode = 0
	v.Success = false
	poolTaobaoCrmMemberJoinurlGetResult.Put(v)
}
