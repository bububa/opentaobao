package bus

import (
	"sync"
)

// TaobaoBusLastplaceGetResult 结构体
type TaobaoBusLastplaceGetResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 目的地数据JSON,具体参考文档说明
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoBusLastplaceGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoBusLastplaceGetResult)
	},
}

// GetTaobaoBusLastplaceGetResult() 从对象池中获取TaobaoBusLastplaceGetResult
func GetTaobaoBusLastplaceGetResult() *TaobaoBusLastplaceGetResult {
	return poolTaobaoBusLastplaceGetResult.Get().(*TaobaoBusLastplaceGetResult)
}

// ReleaseTaobaoBusLastplaceGetResult 释放TaobaoBusLastplaceGetResult
func ReleaseTaobaoBusLastplaceGetResult(v *TaobaoBusLastplaceGetResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = ""
	v.Success = false
	poolTaobaoBusLastplaceGetResult.Put(v)
}
