package servicecenter

import (
	"sync"
)

// TpFundsRecoverResultDo 结构体
type TpFundsRecoverResultDo struct {
	// 实际扣回金额，单位分
	ActualRecoverAmount int64 `json:"actual_recover_amount,omitempty" xml:"actual_recover_amount,omitempty"`
	// 应扣回金额，单位分
	ToRecoverAmount int64 `json:"to_recover_amount,omitempty" xml:"to_recover_amount,omitempty"`
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolTpFundsRecoverResultDo = sync.Pool{
	New: func() any {
		return new(TpFundsRecoverResultDo)
	},
}

// GetTpFundsRecoverResultDo() 从对象池中获取TpFundsRecoverResultDo
func GetTpFundsRecoverResultDo() *TpFundsRecoverResultDo {
	return poolTpFundsRecoverResultDo.Get().(*TpFundsRecoverResultDo)
}

// ReleaseTpFundsRecoverResultDo 释放TpFundsRecoverResultDo
func ReleaseTpFundsRecoverResultDo(v *TpFundsRecoverResultDo) {
	v.ActualRecoverAmount = 0
	v.ToRecoverAmount = 0
	v.BizOrderId = 0
	poolTpFundsRecoverResultDo.Put(v)
}
