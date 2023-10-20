package xhotelitem

import (
	"sync"
)

// TaobaoXhotelItemSelectionSellerStatExposureResult 结构体
type TaobaoXhotelItemSelectionSellerStatExposureResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	Module *TaobaoXhotelItemSelectionSellerStatExposureModule `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelItemSelectionSellerStatExposureResult = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelItemSelectionSellerStatExposureResult)
	},
}

// GetTaobaoXhotelItemSelectionSellerStatExposureResult() 从对象池中获取TaobaoXhotelItemSelectionSellerStatExposureResult
func GetTaobaoXhotelItemSelectionSellerStatExposureResult() *TaobaoXhotelItemSelectionSellerStatExposureResult {
	return poolTaobaoXhotelItemSelectionSellerStatExposureResult.Get().(*TaobaoXhotelItemSelectionSellerStatExposureResult)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatExposureResult 释放TaobaoXhotelItemSelectionSellerStatExposureResult
func ReleaseTaobaoXhotelItemSelectionSellerStatExposureResult(v *TaobaoXhotelItemSelectionSellerStatExposureResult) {
	v.Code = ""
	v.Message = ""
	v.Module = nil
	v.Success = false
	poolTaobaoXhotelItemSelectionSellerStatExposureResult.Put(v)
}
