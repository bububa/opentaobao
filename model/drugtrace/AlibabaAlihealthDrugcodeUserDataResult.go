package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugcodeUserDataResult 结构体
type AlibabaAlihealthDrugcodeUserDataResult struct {
	// 无
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 无
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 无
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 无
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// true 成功  false 不成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugcodeUserDataResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugcodeUserDataResult)
	},
}

// GetAlibabaAlihealthDrugcodeUserDataResult() 从对象池中获取AlibabaAlihealthDrugcodeUserDataResult
func GetAlibabaAlihealthDrugcodeUserDataResult() *AlibabaAlihealthDrugcodeUserDataResult {
	return poolAlibabaAlihealthDrugcodeUserDataResult.Get().(*AlibabaAlihealthDrugcodeUserDataResult)
}

// ReleaseAlibabaAlihealthDrugcodeUserDataResult 释放AlibabaAlihealthDrugcodeUserDataResult
func ReleaseAlibabaAlihealthDrugcodeUserDataResult(v *AlibabaAlihealthDrugcodeUserDataResult) {
	v.Headers = ""
	v.Model = ""
	v.BizExtMap = ""
	v.MappingCode = ""
	v.MsgInfo = ""
	v.MsgCode = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolAlibabaAlihealthDrugcodeUserDataResult.Put(v)
}
