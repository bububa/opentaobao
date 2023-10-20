package alicom

import (
	"sync"
)

// CommonResult 结构体
type CommonResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.Model = ""
	v.Desc = ""
	v.Code = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolCommonResult.Put(v)
}
