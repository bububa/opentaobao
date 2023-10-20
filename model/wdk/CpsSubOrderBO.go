package wdk

import (
	"sync"
)

// CpsSubOrderBO 结构体
type CpsSubOrderBO struct {
	// 业务子订单ID
	BizSubOrderId string `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 商品名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolCpsSubOrderBO = sync.Pool{
	New: func() any {
		return new(CpsSubOrderBO)
	},
}

// GetCpsSubOrderBO() 从对象池中获取CpsSubOrderBO
func GetCpsSubOrderBO() *CpsSubOrderBO {
	return poolCpsSubOrderBO.Get().(*CpsSubOrderBO)
}

// ReleaseCpsSubOrderBO 释放CpsSubOrderBO
func ReleaseCpsSubOrderBO(v *CpsSubOrderBO) {
	v.BizSubOrderId = ""
	v.ItemTitle = ""
	v.Quantity = 0
	poolCpsSubOrderBO.Put(v)
}
