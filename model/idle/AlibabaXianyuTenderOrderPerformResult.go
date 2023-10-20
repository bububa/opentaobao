package idle

import (
	"sync"
)

// AlibabaXianyuTenderOrderPerformResult 结构体
type AlibabaXianyuTenderOrderPerformResult struct {
	// 错误信息
	ErrCodeInfo string `json:"err_code_info,omitempty" xml:"err_code_info,omitempty"`
	// 错误码
	ErrMsgInfo string `json:"err_msg_info,omitempty" xml:"err_msg_info,omitempty"`
	// 请求是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
	// 结果是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlibabaXianyuTenderOrderPerformResult = sync.Pool{
	New: func() any {
		return new(AlibabaXianyuTenderOrderPerformResult)
	},
}

// GetAlibabaXianyuTenderOrderPerformResult() 从对象池中获取AlibabaXianyuTenderOrderPerformResult
func GetAlibabaXianyuTenderOrderPerformResult() *AlibabaXianyuTenderOrderPerformResult {
	return poolAlibabaXianyuTenderOrderPerformResult.Get().(*AlibabaXianyuTenderOrderPerformResult)
}

// ReleaseAlibabaXianyuTenderOrderPerformResult 释放AlibabaXianyuTenderOrderPerformResult
func ReleaseAlibabaXianyuTenderOrderPerformResult(v *AlibabaXianyuTenderOrderPerformResult) {
	v.ErrCodeInfo = ""
	v.ErrMsgInfo = ""
	v.RespSuccess = false
	v.Module = false
	poolAlibabaXianyuTenderOrderPerformResult.Put(v)
}
