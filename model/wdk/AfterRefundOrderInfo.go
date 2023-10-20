package wdk

import (
	"sync"
)

// AfterRefundOrderInfo 结构体
type AfterRefundOrderInfo struct {
	// 退款审核人
	Wdkrc string `json:"wdkrc,omitempty" xml:"wdkrc,omitempty"`
	// 退款审核备注
	Wdkrcm string `json:"wdkrcm,omitempty" xml:"wdkrcm,omitempty"`
	// 退款发起人
	Wdkrfqr string `json:"wdkrfqr,omitempty" xml:"wdkrfqr,omitempty"`
	// 退款理由
	Wdkdfqrrr string `json:"wdkdfqrrr,omitempty" xml:"wdkdfqrrr,omitempty"`
	// 退款发起备注
	Wdkrfqrmemo string `json:"wdkrfqrmemo,omitempty" xml:"wdkrfqrmemo,omitempty"`
	// 外部退款单号
	Rprid string `json:"rprid,omitempty" xml:"rprid,omitempty"`
	// 实际退货数量
	Wdkfra string `json:"wdkfra,omitempty" xml:"wdkfra,omitempty"`
	// 实际退款金额
	Wdkrf string `json:"wdkrf,omitempty" xml:"wdkrf,omitempty"`
}

var poolAfterRefundOrderInfo = sync.Pool{
	New: func() any {
		return new(AfterRefundOrderInfo)
	},
}

// GetAfterRefundOrderInfo() 从对象池中获取AfterRefundOrderInfo
func GetAfterRefundOrderInfo() *AfterRefundOrderInfo {
	return poolAfterRefundOrderInfo.Get().(*AfterRefundOrderInfo)
}

// ReleaseAfterRefundOrderInfo 释放AfterRefundOrderInfo
func ReleaseAfterRefundOrderInfo(v *AfterRefundOrderInfo) {
	v.Wdkrc = ""
	v.Wdkrcm = ""
	v.Wdkrfqr = ""
	v.Wdkdfqrrr = ""
	v.Wdkrfqrmemo = ""
	v.Rprid = ""
	v.Wdkfra = ""
	v.Wdkrf = ""
	poolAfterRefundOrderInfo.Put(v)
}
