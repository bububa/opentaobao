package alsc

import (
	"sync"
)

// CommonResult 结构体
type CommonResult struct {
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 错误结果显示
	ResultView string `json:"result_view,omitempty" xml:"result_view,omitempty"`
	// 成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
	// 开卡是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
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
	v.ResultCode = ""
	v.ResultDesc = ""
	v.ResultView = ""
	v.BizSuccess = false
	v.Result = false
	poolCommonResult.Put(v)
}
