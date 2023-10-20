package tbk

import (
	"sync"
)

// TaobaoTbkCouponConvertRpcResult 结构体
type TaobaoTbkCouponConvertRpcResult struct {
	// 见错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// data
	Results *MaterialDto `json:"results,omitempty" xml:"results,omitempty"`
	// 见错误码描述
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoTbkCouponConvertRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkCouponConvertRpcResult)
	},
}

// GetTaobaoTbkCouponConvertRpcResult() 从对象池中获取TaobaoTbkCouponConvertRpcResult
func GetTaobaoTbkCouponConvertRpcResult() *TaobaoTbkCouponConvertRpcResult {
	return poolTaobaoTbkCouponConvertRpcResult.Get().(*TaobaoTbkCouponConvertRpcResult)
}

// ReleaseTaobaoTbkCouponConvertRpcResult 释放TaobaoTbkCouponConvertRpcResult
func ReleaseTaobaoTbkCouponConvertRpcResult(v *TaobaoTbkCouponConvertRpcResult) {
	v.ResultMsg = ""
	v.Results = nil
	v.ResultCode = 0
	poolTaobaoTbkCouponConvertRpcResult.Put(v)
}
