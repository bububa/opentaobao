package alihealth2

import (
	"sync"
)

// TaobaoTradeDrugOrderGetResultSet 结构体
type TaobaoTradeDrugOrderGetResultSet struct {
	// 异常信息
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// TakeoutThirdOrder订单数据
	Result *TakeoutThirdOrder `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoTradeDrugOrderGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoTradeDrugOrderGetResultSet)
	},
}

// GetTaobaoTradeDrugOrderGetResultSet() 从对象池中获取TaobaoTradeDrugOrderGetResultSet
func GetTaobaoTradeDrugOrderGetResultSet() *TaobaoTradeDrugOrderGetResultSet {
	return poolTaobaoTradeDrugOrderGetResultSet.Get().(*TaobaoTradeDrugOrderGetResultSet)
}

// ReleaseTaobaoTradeDrugOrderGetResultSet 释放TaobaoTradeDrugOrderGetResultSet
func ReleaseTaobaoTradeDrugOrderGetResultSet(v *TaobaoTradeDrugOrderGetResultSet) {
	v.Exception = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	poolTaobaoTradeDrugOrderGetResultSet.Put(v)
}
