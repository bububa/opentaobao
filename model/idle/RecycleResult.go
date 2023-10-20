package idle

import (
	"sync"
)

// RecycleResult 结构体
type RecycleResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 闲鱼状态 0：为在线 1：待测试
	SpuStatus int64 `json:"spu_status,omitempty" xml:"spu_status,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRecycleResult = sync.Pool{
	New: func() any {
		return new(RecycleResult)
	},
}

// GetRecycleResult() 从对象池中获取RecycleResult
func GetRecycleResult() *RecycleResult {
	return poolRecycleResult.Get().(*RecycleResult)
}

// ReleaseRecycleResult 释放RecycleResult
func ReleaseRecycleResult(v *RecycleResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.ErrMessage = ""
	v.SpuStatus = 0
	v.Success = false
	poolRecycleResult.Put(v)
}
