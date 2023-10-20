package moscm

import (
	"sync"
)

// InputCspuResult 结构体
type InputCspuResult struct {
	// 中台商品id
	CspuId string `json:"cspu_id,omitempty" xml:"cspu_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 供应商商品id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 是否录入成功，true：成功 false：失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolInputCspuResult = sync.Pool{
	New: func() any {
		return new(InputCspuResult)
	},
}

// GetInputCspuResult() 从对象池中获取InputCspuResult
func GetInputCspuResult() *InputCspuResult {
	return poolInputCspuResult.Get().(*InputCspuResult)
}

// ReleaseInputCspuResult 释放InputCspuResult
func ReleaseInputCspuResult(v *InputCspuResult) {
	v.CspuId = ""
	v.Message = ""
	v.OuterId = ""
	v.Success = false
	poolInputCspuResult.Put(v)
}
