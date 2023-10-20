package idle

import (
	"sync"
)

// AlibabaIdleConsignmentiiOrderGetResult 结构体
type AlibabaIdleConsignmentiiOrderGetResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单详情
	Module *ConsignmentV2OrderTO `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleConsignmentiiOrderGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentiiOrderGetResult)
	},
}

// GetAlibabaIdleConsignmentiiOrderGetResult() 从对象池中获取AlibabaIdleConsignmentiiOrderGetResult
func GetAlibabaIdleConsignmentiiOrderGetResult() *AlibabaIdleConsignmentiiOrderGetResult {
	return poolAlibabaIdleConsignmentiiOrderGetResult.Get().(*AlibabaIdleConsignmentiiOrderGetResult)
}

// ReleaseAlibabaIdleConsignmentiiOrderGetResult 释放AlibabaIdleConsignmentiiOrderGetResult
func ReleaseAlibabaIdleConsignmentiiOrderGetResult(v *AlibabaIdleConsignmentiiOrderGetResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleConsignmentiiOrderGetResult.Put(v)
}
