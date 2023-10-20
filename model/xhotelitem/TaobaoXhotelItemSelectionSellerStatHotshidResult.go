package xhotelitem

import (
	"sync"
)

// TaobaoXhotelItemSelectionSellerStatHotshidResult 结构体
type TaobaoXhotelItemSelectionSellerStatHotshidResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结构
	Module *TaobaoXhotelItemSelectionSellerStatHotshidModule `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelItemSelectionSellerStatHotshidResult = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelItemSelectionSellerStatHotshidResult)
	},
}

// GetTaobaoXhotelItemSelectionSellerStatHotshidResult() 从对象池中获取TaobaoXhotelItemSelectionSellerStatHotshidResult
func GetTaobaoXhotelItemSelectionSellerStatHotshidResult() *TaobaoXhotelItemSelectionSellerStatHotshidResult {
	return poolTaobaoXhotelItemSelectionSellerStatHotshidResult.Get().(*TaobaoXhotelItemSelectionSellerStatHotshidResult)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatHotshidResult 释放TaobaoXhotelItemSelectionSellerStatHotshidResult
func ReleaseTaobaoXhotelItemSelectionSellerStatHotshidResult(v *TaobaoXhotelItemSelectionSellerStatHotshidResult) {
	v.Code = ""
	v.Message = ""
	v.Module = nil
	v.Success = false
	poolTaobaoXhotelItemSelectionSellerStatHotshidResult.Put(v)
}
