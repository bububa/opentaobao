package product

import (
	"sync"
)

// TaobaoBanamadpcItemSelectPropApiResult 结构体
type TaobaoBanamadpcItemSelectPropApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 入参类目下入参属性的子属性schema xml
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoBanamadpcItemSelectPropApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemSelectPropApiResult)
	},
}

// GetTaobaoBanamadpcItemSelectPropApiResult() 从对象池中获取TaobaoBanamadpcItemSelectPropApiResult
func GetTaobaoBanamadpcItemSelectPropApiResult() *TaobaoBanamadpcItemSelectPropApiResult {
	return poolTaobaoBanamadpcItemSelectPropApiResult.Get().(*TaobaoBanamadpcItemSelectPropApiResult)
}

// ReleaseTaobaoBanamadpcItemSelectPropApiResult 释放TaobaoBanamadpcItemSelectPropApiResult
func ReleaseTaobaoBanamadpcItemSelectPropApiResult(v *TaobaoBanamadpcItemSelectPropApiResult) {
	v.ErMsg = ""
	v.ErCode = ""
	v.Result = ""
	v.Error = false
	poolTaobaoBanamadpcItemSelectPropApiResult.Put(v)
}
