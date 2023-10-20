package wdk

import (
	"sync"
)

// TxdBillListGetResult 结构体
type TxdBillListGetResult struct {
	// 返回值对象列表
	TxdBillDetailBOS []TxdBillDetailBo `json:"txd_bill_detail_b_o_s,omitempty" xml:"txd_bill_detail_b_o_s>txd_bill_detail_bo,omitempty"`
	// 结果总调条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolTxdBillListGetResult = sync.Pool{
	New: func() any {
		return new(TxdBillListGetResult)
	},
}

// GetTxdBillListGetResult() 从对象池中获取TxdBillListGetResult
func GetTxdBillListGetResult() *TxdBillListGetResult {
	return poolTxdBillListGetResult.Get().(*TxdBillListGetResult)
}

// ReleaseTxdBillListGetResult 释放TxdBillListGetResult
func ReleaseTxdBillListGetResult(v *TxdBillListGetResult) {
	v.TxdBillDetailBOS = v.TxdBillDetailBOS[:0]
	v.Total = 0
	poolTxdBillListGetResult.Put(v)
}
