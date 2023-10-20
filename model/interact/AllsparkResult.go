package interact

import (
	"sync"
)

// AllsparkResult 结构体
type AllsparkResult struct {
	// 出错提示
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 出错代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 服务结果对象
	Data *ActivityWriteResult `json:"data,omitempty" xml:"data,omitempty"`
	// 是否注册成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAllsparkResult = sync.Pool{
	New: func() any {
		return new(AllsparkResult)
	},
}

// GetAllsparkResult() 从对象池中获取AllsparkResult
func GetAllsparkResult() *AllsparkResult {
	return poolAllsparkResult.Get().(*AllsparkResult)
}

// ReleaseAllsparkResult 释放AllsparkResult
func ReleaseAllsparkResult(v *AllsparkResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Data = nil
	v.Success = false
	poolAllsparkResult.Put(v)
}
