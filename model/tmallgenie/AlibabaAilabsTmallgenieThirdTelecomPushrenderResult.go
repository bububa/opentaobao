package tmallgenie

import (
	"sync"
)

// AlibabaAilabsTmallgenieThirdTelecomPushrenderResult 结构体
type AlibabaAilabsTmallgenieThirdTelecomPushrenderResult struct {
	// 结果信息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 状态码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAilabsTmallgenieThirdTelecomPushrenderResult = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieThirdTelecomPushrenderResult)
	},
}

// GetAlibabaAilabsTmallgenieThirdTelecomPushrenderResult() 从对象池中获取AlibabaAilabsTmallgenieThirdTelecomPushrenderResult
func GetAlibabaAilabsTmallgenieThirdTelecomPushrenderResult() *AlibabaAilabsTmallgenieThirdTelecomPushrenderResult {
	return poolAlibabaAilabsTmallgenieThirdTelecomPushrenderResult.Get().(*AlibabaAilabsTmallgenieThirdTelecomPushrenderResult)
}

// ReleaseAlibabaAilabsTmallgenieThirdTelecomPushrenderResult 释放AlibabaAilabsTmallgenieThirdTelecomPushrenderResult
func ReleaseAlibabaAilabsTmallgenieThirdTelecomPushrenderResult(v *AlibabaAilabsTmallgenieThirdTelecomPushrenderResult) {
	v.ResultMessage = ""
	v.ResultCode = 0
	v.Data = false
	poolAlibabaAilabsTmallgenieThirdTelecomPushrenderResult.Put(v)
}
