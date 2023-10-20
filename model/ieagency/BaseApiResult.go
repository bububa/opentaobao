package ieagency

import (
	"sync"
)

// BaseApiResult 结构体
type BaseApiResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误编码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 创建支付订单结果
	Model *IeBookPayOrderVo `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBaseApiResult = sync.Pool{
	New: func() any {
		return new(BaseApiResult)
	},
}

// GetBaseApiResult() 从对象池中获取BaseApiResult
func GetBaseApiResult() *BaseApiResult {
	return poolBaseApiResult.Get().(*BaseApiResult)
}

// ReleaseBaseApiResult 释放BaseApiResult
func ReleaseBaseApiResult(v *BaseApiResult) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Model = nil
	v.Success = false
	poolBaseApiResult.Put(v)
}
