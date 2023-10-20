package tmallgenie

import (
	"sync"
)

// AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult 结构体
type AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult struct {
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult)
	},
}

// GetAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult() 从对象池中获取AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult
func GetAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult() *AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult {
	return poolAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult.Get().(*AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult)
}

// ReleaseAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult 释放AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult
func ReleaseAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult(v *AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult) {
	v.ResultCode = ""
	v.Message = ""
	poolAlibabaAilabsTmallgenieThirdUnicomShenyanOperResult.Put(v)
}
