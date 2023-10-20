package xhotelitem

import (
	"sync"
)

// TaobaoXhotelBnbpromoGetResultSet 结构体
type TaobaoXhotelBnbpromoGetResultSet struct {
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 民宿活动信息
	BnbPromo *BnbPromoDto `json:"bnb_promo,omitempty" xml:"bnb_promo,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelBnbpromoGetResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoGetResultSet)
	},
}

// GetTaobaoXhotelBnbpromoGetResultSet() 从对象池中获取TaobaoXhotelBnbpromoGetResultSet
func GetTaobaoXhotelBnbpromoGetResultSet() *TaobaoXhotelBnbpromoGetResultSet {
	return poolTaobaoXhotelBnbpromoGetResultSet.Get().(*TaobaoXhotelBnbpromoGetResultSet)
}

// ReleaseTaobaoXhotelBnbpromoGetResultSet 释放TaobaoXhotelBnbpromoGetResultSet
func ReleaseTaobaoXhotelBnbpromoGetResultSet(v *TaobaoXhotelBnbpromoGetResultSet) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.BnbPromo = nil
	v.Success = false
	poolTaobaoXhotelBnbpromoGetResultSet.Put(v)
}
