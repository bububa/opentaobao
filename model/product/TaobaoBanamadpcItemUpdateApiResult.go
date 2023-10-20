package product

import (
	"sync"
)

// TaobaoBanamadpcItemUpdateApiResult 结构体
type TaobaoBanamadpcItemUpdateApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 商品id
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoBanamadpcItemUpdateApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemUpdateApiResult)
	},
}

// GetTaobaoBanamadpcItemUpdateApiResult() 从对象池中获取TaobaoBanamadpcItemUpdateApiResult
func GetTaobaoBanamadpcItemUpdateApiResult() *TaobaoBanamadpcItemUpdateApiResult {
	return poolTaobaoBanamadpcItemUpdateApiResult.Get().(*TaobaoBanamadpcItemUpdateApiResult)
}

// ReleaseTaobaoBanamadpcItemUpdateApiResult 释放TaobaoBanamadpcItemUpdateApiResult
func ReleaseTaobaoBanamadpcItemUpdateApiResult(v *TaobaoBanamadpcItemUpdateApiResult) {
	v.ErMsg = ""
	v.ErCode = ""
	v.Result = 0
	v.Error = false
	poolTaobaoBanamadpcItemUpdateApiResult.Put(v)
}
