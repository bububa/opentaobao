package ascp

import (
	"sync"
)

// ConfirmOrderLines 结构体
type ConfirmOrderLines struct {
	// 交易主单号
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 交易子单号
	SubSourceOrderCode string `json:"sub_source_order_code,omitempty" xml:"sub_source_order_code,omitempty"`
	// 单据行号，与order_lines中的order_line_no需要对应
	OrderLineNo string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
	// ERP子货品编码（可以跟接单时不一致）
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 实发数量
	ActualQty int64 `json:"actual_qty,omitempty" xml:"actual_qty,omitempty"`
}

var poolConfirmOrderLines = sync.Pool{
	New: func() any {
		return new(ConfirmOrderLines)
	},
}

// GetConfirmOrderLines() 从对象池中获取ConfirmOrderLines
func GetConfirmOrderLines() *ConfirmOrderLines {
	return poolConfirmOrderLines.Get().(*ConfirmOrderLines)
}

// ReleaseConfirmOrderLines 释放ConfirmOrderLines
func ReleaseConfirmOrderLines(v *ConfirmOrderLines) {
	v.SourceOrderCode = ""
	v.SubSourceOrderCode = ""
	v.OrderLineNo = ""
	v.ScItemId = ""
	v.ActualQty = 0
	poolConfirmOrderLines.Put(v)
}
