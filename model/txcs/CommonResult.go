package txcs

import (
	"sync"
)

// CommonResult 结构体
type CommonResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 创建的互动实例
	Data *StatementBillDto `json:"data,omitempty" xml:"data,omitempty"`
	// 服务标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 服务标识
	BizException bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
}

var poolCommonResult = sync.Pool{
	New: func() any {
		return new(CommonResult)
	},
}

// GetCommonResult() 从对象池中获取CommonResult
func GetCommonResult() *CommonResult {
	return poolCommonResult.Get().(*CommonResult)
}

// ReleaseCommonResult 释放CommonResult
func ReleaseCommonResult(v *CommonResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	v.BizException = false
	poolCommonResult.Put(v)
}
