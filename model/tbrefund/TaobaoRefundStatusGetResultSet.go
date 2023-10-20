package tbrefund

import (
	"sync"
)

// TaobaoRefundStatusGetResultSet 结构体
type TaobaoRefundStatusGetResultSet struct {
	// 数组对象
	ResultList []QueryRefundStatusResponse `json:"result_list,omitempty" xml:"result_list>query_refund_status_response,omitempty"`
	// 错误码，没有表示无异常
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

var poolTaobaoRefundStatusGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoRefundStatusGetResultSet)
	},
}

// GetTaobaoRefundStatusGetResultSet() 从对象池中获取TaobaoRefundStatusGetResultSet
func GetTaobaoRefundStatusGetResultSet() *TaobaoRefundStatusGetResultSet {
	return poolTaobaoRefundStatusGetResultSet.Get().(*TaobaoRefundStatusGetResultSet)
}

// ReleaseTaobaoRefundStatusGetResultSet 释放TaobaoRefundStatusGetResultSet
func ReleaseTaobaoRefundStatusGetResultSet(v *TaobaoRefundStatusGetResultSet) {
	v.ResultList = v.ResultList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	poolTaobaoRefundStatusGetResultSet.Put(v)
}
