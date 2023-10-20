package idle

import (
	"sync"
)

// AlibabaIdleSpuRegisterModifyResult 结构体
type AlibabaIdleSpuRegisterModifyResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// spu挂载信息后台状态，0在线，-1删除，1测试中。当spu第一次挂载时，会进入1测试中状态。服务商联调通过后，需要再次挂载，actionType还传0，挂载信息状态会变成0已上线。
	Module int64 `json:"module,omitempty" xml:"module,omitempty"`
	// 是否挂载成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleSpuRegisterModifyResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleSpuRegisterModifyResult)
	},
}

// GetAlibabaIdleSpuRegisterModifyResult() 从对象池中获取AlibabaIdleSpuRegisterModifyResult
func GetAlibabaIdleSpuRegisterModifyResult() *AlibabaIdleSpuRegisterModifyResult {
	return poolAlibabaIdleSpuRegisterModifyResult.Get().(*AlibabaIdleSpuRegisterModifyResult)
}

// ReleaseAlibabaIdleSpuRegisterModifyResult 释放AlibabaIdleSpuRegisterModifyResult
func ReleaseAlibabaIdleSpuRegisterModifyResult(v *AlibabaIdleSpuRegisterModifyResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = 0
	v.Success = false
	poolAlibabaIdleSpuRegisterModifyResult.Put(v)
}
