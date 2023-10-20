package wdk

import (
	"sync"
)

// TopBaseResult 结构体
type TopBaseResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误码备注说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 订单对象
	Model *OrderResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopBaseResult = sync.Pool{
	New: func() any {
		return new(TopBaseResult)
	},
}

// GetTopBaseResult() 从对象池中获取TopBaseResult
func GetTopBaseResult() *TopBaseResult {
	return poolTopBaseResult.Get().(*TopBaseResult)
}

// ReleaseTopBaseResult 释放TopBaseResult
func ReleaseTopBaseResult(v *TopBaseResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.Model = nil
	v.Success = false
	poolTopBaseResult.Put(v)
}
