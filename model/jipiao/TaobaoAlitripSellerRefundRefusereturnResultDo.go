package jipiao

import (
	"sync"
)

// TaobaoAlitripSellerRefundRefusereturnResultDo 结构体
type TaobaoAlitripSellerRefundRefusereturnResultDo struct {
	// 系统自动生成
	ErrorCode string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 系统自动生成
	ErrorMsg string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 处理结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAlitripSellerRefundRefusereturnResultDo = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripSellerRefundRefusereturnResultDo)
	},
}

// GetTaobaoAlitripSellerRefundRefusereturnResultDo() 从对象池中获取TaobaoAlitripSellerRefundRefusereturnResultDo
func GetTaobaoAlitripSellerRefundRefusereturnResultDo() *TaobaoAlitripSellerRefundRefusereturnResultDo {
	return poolTaobaoAlitripSellerRefundRefusereturnResultDo.Get().(*TaobaoAlitripSellerRefundRefusereturnResultDo)
}

// ReleaseTaobaoAlitripSellerRefundRefusereturnResultDo 释放TaobaoAlitripSellerRefundRefusereturnResultDo
func ReleaseTaobaoAlitripSellerRefundRefusereturnResultDo(v *TaobaoAlitripSellerRefundRefusereturnResultDo) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = false
	v.Success = false
	poolTaobaoAlitripSellerRefundRefusereturnResultDo.Put(v)
}
