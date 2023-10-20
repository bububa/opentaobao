package jipiao

import (
	"sync"
)

// TaobaoAlitripSellerRefundGetResultDo 结构体
type TaobaoAlitripSellerRefundGetResultDo struct {
	// 系统自动生成
	ErrorCode string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 申请单详情
	Results *ReturnTicketDetail `json:"results,omitempty" xml:"results,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripSellerRefundGetResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundGetResultDo)
	},
}

// GetTaobaoAlitripSellerRefundGetResultDo() 从对象池中获取TaobaoAlitripSellerRefundGetResultDo
func GetTaobaoAlitripSellerRefundGetResultDo() *TaobaoAlitripSellerRefundGetResultDo {
	return poolTaobaoAlitripSellerRefundGetResultDo.Get().(*TaobaoAlitripSellerRefundGetResultDo)
}

// ReleaseTaobaoAlitripSellerRefundGetResultDo 释放TaobaoAlitripSellerRefundGetResultDo
func ReleaseTaobaoAlitripSellerRefundGetResultDo(v *TaobaoAlitripSellerRefundGetResultDo) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Results = nil
	v.Success = false
	poolTaobaoAlitripSellerRefundGetResultDo.Put(v)
}
