package tmallhk

import (
	"sync"
)

// ClearanceDataDo 结构体
type ClearanceDataDo struct {
	// 订单数据封装
	BizOrderDO *ClearanceLogisticsOrderDo `json:"biz_order_d_o,omitempty" xml:"biz_order_d_o,omitempty"`
	// 支付单封装
	PayOrderDO *ClearancePayOrderDo `json:"pay_order_d_o,omitempty" xml:"pay_order_d_o,omitempty"`
}

var poolClearanceDataDo = sync.Pool{
	New: func() any {
		return new(ClearanceDataDo)
	},
}

// GetClearanceDataDo() 从对象池中获取ClearanceDataDo
func GetClearanceDataDo() *ClearanceDataDo {
	return poolClearanceDataDo.Get().(*ClearanceDataDo)
}

// ReleaseClearanceDataDo 释放ClearanceDataDo
func ReleaseClearanceDataDo(v *ClearanceDataDo) {
	v.BizOrderDO = nil
	v.PayOrderDO = nil
	poolClearanceDataDo.Put(v)
}
