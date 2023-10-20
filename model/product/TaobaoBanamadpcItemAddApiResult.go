package product

import (
	"sync"
)

// TaobaoBanamadpcItemAddApiResult 结构体
type TaobaoBanamadpcItemAddApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 商品id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoBanamadpcItemAddApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemAddApiResult)
	},
}

// GetTaobaoBanamadpcItemAddApiResult() 从对象池中获取TaobaoBanamadpcItemAddApiResult
func GetTaobaoBanamadpcItemAddApiResult() *TaobaoBanamadpcItemAddApiResult {
	return poolTaobaoBanamadpcItemAddApiResult.Get().(*TaobaoBanamadpcItemAddApiResult)
}

// ReleaseTaobaoBanamadpcItemAddApiResult 释放TaobaoBanamadpcItemAddApiResult
func ReleaseTaobaoBanamadpcItemAddApiResult(v *TaobaoBanamadpcItemAddApiResult) {
	v.ErMsg = ""
	v.ErCode = ""
	v.Result = 0
	v.Error = false
	poolTaobaoBanamadpcItemAddApiResult.Put(v)
}
