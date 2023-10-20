package cntms

import (
	"sync"
)

// CnTmsLogisticsOrderItem 结构体
type CnTmsLogisticsOrderItem struct {
	// ERP订单明细编码
	OrderItemId string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	// 子交易号
	SubTradeId string `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 发货商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 扩展字段 K:V;
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 发货商品商品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品单价，单位分
	ItemPrice int64 `json:"item_price,omitempty" xml:"item_price,omitempty"`
}

var poolCnTmsLogisticsOrderItem = sync.Pool{
	New: func() any {
		return new(CnTmsLogisticsOrderItem)
	},
}

// GetCnTmsLogisticsOrderItem() 从对象池中获取CnTmsLogisticsOrderItem
func GetCnTmsLogisticsOrderItem() *CnTmsLogisticsOrderItem {
	return poolCnTmsLogisticsOrderItem.Get().(*CnTmsLogisticsOrderItem)
}

// ReleaseCnTmsLogisticsOrderItem 释放CnTmsLogisticsOrderItem
func ReleaseCnTmsLogisticsOrderItem(v *CnTmsLogisticsOrderItem) {
	v.OrderItemId = ""
	v.SubTradeId = ""
	v.ItemName = ""
	v.ExtendFields = ""
	v.Remark = ""
	v.Quantity = 0
	v.ItemPrice = 0
	poolCnTmsLogisticsOrderItem.Put(v)
}
