package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotUseroscontrolResult 结构体
type AlibabaAliqinFcIotUseroscontrolResult struct {
	// 停开结果描述
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 停开结果编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotUseroscontrolResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotUseroscontrolResult)
	},
}

// GetAlibabaAliqinFcIotUseroscontrolResult() 从对象池中获取AlibabaAliqinFcIotUseroscontrolResult
func GetAlibabaAliqinFcIotUseroscontrolResult() *AlibabaAliqinFcIotUseroscontrolResult {
	return poolAlibabaAliqinFcIotUseroscontrolResult.Get().(*AlibabaAliqinFcIotUseroscontrolResult)
}

// ReleaseAlibabaAliqinFcIotUseroscontrolResult 释放AlibabaAliqinFcIotUseroscontrolResult
func ReleaseAlibabaAliqinFcIotUseroscontrolResult(v *AlibabaAliqinFcIotUseroscontrolResult) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIotUseroscontrolResult.Put(v)
}
