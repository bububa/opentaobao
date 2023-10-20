package wdk

import (
	"sync"
)

// OrderOperateResult 结构体
type OrderOperateResult struct {
	// 子单列表信息
	SubOrderResults []SubOrder `json:"sub_order_results,omitempty" xml:"sub_order_results>sub_order,omitempty"`
	// 盒马主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

var poolOrderOperateResult = sync.Pool{
	New: func() any {
		return new(OrderOperateResult)
	},
}

// GetOrderOperateResult() 从对象池中获取OrderOperateResult
func GetOrderOperateResult() *OrderOperateResult {
	return poolOrderOperateResult.Get().(*OrderOperateResult)
}

// ReleaseOrderOperateResult 释放OrderOperateResult
func ReleaseOrderOperateResult(v *OrderOperateResult) {
	v.SubOrderResults = v.SubOrderResults[:0]
	v.BizOrderId = ""
	v.OutOrderId = ""
	poolOrderOperateResult.Put(v)
}
