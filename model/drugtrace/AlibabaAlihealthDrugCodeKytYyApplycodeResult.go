package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugCodeKytYyApplycodeResult 结构体
type AlibabaAlihealthDrugCodeKytYyApplycodeResult struct {
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 父子码关系对象
	Model *AlibabaAlihealthDrugCodeKytYyApplycodeModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugCodeKytYyApplycodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugCodeKytYyApplycodeResult)
	},
}

// GetAlibabaAlihealthDrugCodeKytYyApplycodeResult() 从对象池中获取AlibabaAlihealthDrugCodeKytYyApplycodeResult
func GetAlibabaAlihealthDrugCodeKytYyApplycodeResult() *AlibabaAlihealthDrugCodeKytYyApplycodeResult {
	return poolAlibabaAlihealthDrugCodeKytYyApplycodeResult.Get().(*AlibabaAlihealthDrugCodeKytYyApplycodeResult)
}

// ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeResult 释放AlibabaAlihealthDrugCodeKytYyApplycodeResult
func ReleaseAlibabaAlihealthDrugCodeKytYyApplycodeResult(v *AlibabaAlihealthDrugCodeKytYyApplycodeResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugCodeKytYyApplycodeResult.Put(v)
}
