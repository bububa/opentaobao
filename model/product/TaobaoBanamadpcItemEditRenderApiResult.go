package product

import (
	"sync"
)

// TaobaoBanamadpcItemEditRenderApiResult 结构体
type TaobaoBanamadpcItemEditRenderApiResult struct {
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 错误码
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 商品标题，价格，图片等商品数据的schema xml
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

var poolTaobaoBanamadpcItemEditRenderApiResult = sync.Pool{
	New: func() any {
		return new(TaobaoBanamadpcItemEditRenderApiResult)
	},
}

// GetTaobaoBanamadpcItemEditRenderApiResult() 从对象池中获取TaobaoBanamadpcItemEditRenderApiResult
func GetTaobaoBanamadpcItemEditRenderApiResult() *TaobaoBanamadpcItemEditRenderApiResult {
	return poolTaobaoBanamadpcItemEditRenderApiResult.Get().(*TaobaoBanamadpcItemEditRenderApiResult)
}

// ReleaseTaobaoBanamadpcItemEditRenderApiResult 释放TaobaoBanamadpcItemEditRenderApiResult
func ReleaseTaobaoBanamadpcItemEditRenderApiResult(v *TaobaoBanamadpcItemEditRenderApiResult) {
	v.ErMsg = ""
	v.ErCode = ""
	v.Result = ""
	v.Error = false
	poolTaobaoBanamadpcItemEditRenderApiResult.Put(v)
}
