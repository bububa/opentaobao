package jipiao

import (
	"sync"
)

// TaobaoAlitripSellerRefundSearchResultDo 结构体
type TaobaoAlitripSellerRefundSearchResultDo struct {
	// ReturnTicketDo
	Results []ReturnTicketDo `json:"results,omitempty" xml:"results>return_ticket_do,omitempty"`
	// 错误码
	ErrorCode string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripSellerRefundSearchResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundSearchResultDo)
	},
}

// GetTaobaoAlitripSellerRefundSearchResultDo() 从对象池中获取TaobaoAlitripSellerRefundSearchResultDo
func GetTaobaoAlitripSellerRefundSearchResultDo() *TaobaoAlitripSellerRefundSearchResultDo {
	return poolTaobaoAlitripSellerRefundSearchResultDo.Get().(*TaobaoAlitripSellerRefundSearchResultDo)
}

// ReleaseTaobaoAlitripSellerRefundSearchResultDo 释放TaobaoAlitripSellerRefundSearchResultDo
func ReleaseTaobaoAlitripSellerRefundSearchResultDo(v *TaobaoAlitripSellerRefundSearchResultDo) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoAlitripSellerRefundSearchResultDo.Put(v)
}
