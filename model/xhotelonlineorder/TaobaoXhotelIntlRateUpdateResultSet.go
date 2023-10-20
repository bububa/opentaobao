package xhotelonlineorder

import (
	"sync"
)

// TaobaoXhotelIntlRateUpdateResultSet 结构体
type TaobaoXhotelIntlRateUpdateResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// sub_error_msg
	SubErrorMsg string `json:"sub_error_msg,omitempty" xml:"sub_error_msg,omitempty"`
	// sub_error_code
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// detail_msg_info
	DetailMsgInfo string `json:"detail_msg_info,omitempty" xml:"detail_msg_info,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// total_results
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 总记录条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolTaobaoXhotelIntlRateUpdateResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelIntlRateUpdateResultSet)
	},
}

// GetTaobaoXhotelIntlRateUpdateResultSet() 从对象池中获取TaobaoXhotelIntlRateUpdateResultSet
func GetTaobaoXhotelIntlRateUpdateResultSet() *TaobaoXhotelIntlRateUpdateResultSet {
	return poolTaobaoXhotelIntlRateUpdateResultSet.Get().(*TaobaoXhotelIntlRateUpdateResultSet)
}

// ReleaseTaobaoXhotelIntlRateUpdateResultSet 释放TaobaoXhotelIntlRateUpdateResultSet
func ReleaseTaobaoXhotelIntlRateUpdateResultSet(v *TaobaoXhotelIntlRateUpdateResultSet) {
	v.ErrorCode = ""
	v.SubErrorMsg = ""
	v.SubErrorCode = ""
	v.DetailMsgInfo = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.TotalCount = 0
	poolTaobaoXhotelIntlRateUpdateResultSet.Put(v)
}
