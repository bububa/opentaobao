package ascp

import (
	"sync"
)

// ScItems 结构体
type ScItems struct {
	// 单据行号，与order_lines中的order_line_no需要对应
	OrderLineNo string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
	// 子件运单号，如果有子母件可填
	SubExpressCode string `json:"sub_express_code,omitempty" xml:"sub_express_code,omitempty"`
	// 包裹内ERP子货品编码
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 包裹内实发货品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolScItems = sync.Pool{
	New: func() any {
		return new(ScItems)
	},
}

// GetScItems() 从对象池中获取ScItems
func GetScItems() *ScItems {
	return poolScItems.Get().(*ScItems)
}

// ReleaseScItems 释放ScItems
func ReleaseScItems(v *ScItems) {
	v.OrderLineNo = ""
	v.SubExpressCode = ""
	v.ScItemId = ""
	v.Quantity = 0
	poolScItems.Put(v)
}
