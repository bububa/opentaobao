package wlbimports

import (
	"sync"
)

// TaobaoWlbCrossborderWaybillGetTopResult 结构体
type TaobaoWlbCrossborderWaybillGetTopResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 子错误信息
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// 子错误码
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// 面单信息
	Result *WayBillDto `json:"result,omitempty" xml:"result,omitempty"`
	// 结果数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWlbCrossborderWaybillGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbCrossborderWaybillGetTopResult)
	},
}

// GetTaobaoWlbCrossborderWaybillGetTopResult() 从对象池中获取TaobaoWlbCrossborderWaybillGetTopResult
func GetTaobaoWlbCrossborderWaybillGetTopResult() *TaobaoWlbCrossborderWaybillGetTopResult {
	return poolTaobaoWlbCrossborderWaybillGetTopResult.Get().(*TaobaoWlbCrossborderWaybillGetTopResult)
}

// ReleaseTaobaoWlbCrossborderWaybillGetTopResult 释放TaobaoWlbCrossborderWaybillGetTopResult
func ReleaseTaobaoWlbCrossborderWaybillGetTopResult(v *TaobaoWlbCrossborderWaybillGetTopResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.SubErrorCode = ""
	v.SubErrorMsg = ""
	v.Result = nil
	v.TotalResults = 0
	v.Success = false
	poolTaobaoWlbCrossborderWaybillGetTopResult.Put(v)
}
