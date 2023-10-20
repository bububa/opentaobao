package pur

import (
	"sync"
)

// SupplierPreInvoiceItem 结构体
type SupplierPreInvoiceItem struct {
	// 关闭原因
	LineCloseReason string `json:"line_close_reason,omitempty" xml:"line_close_reason,omitempty"`
	// 行状态
	LineStatus string `json:"line_status,omitempty" xml:"line_status,omitempty"`
	// 关联开票单行号
	SourceKpLineNo string `json:"source_kp_line_no,omitempty" xml:"source_kp_line_no,omitempty"`
	// 关联开票单单号
	SourceKpNo string `json:"source_kp_no,omitempty" xml:"source_kp_no,omitempty"`
	// 应付发票行备注
	LineRemark string `json:"line_remark,omitempty" xml:"line_remark,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 税码
	TaxCode string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
	// 发票行税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 发票行总金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 发票行数量
	InvoiceLineQuantity string `json:"invoice_line_quantity,omitempty" xml:"invoice_line_quantity,omitempty"`
	// 规格型号
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 货物或服务名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 货物或服务id
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 匹配入库id
	SourceInventory string `json:"source_inventory,omitempty" xml:"source_inventory,omitempty"`
	// 匹配接收行编号
	SourceRtLineNo string `json:"source_rt_line_no,omitempty" xml:"source_rt_line_no,omitempty"`
	// 匹配接收行编号
	SourceRtNo string `json:"source_rt_no,omitempty" xml:"source_rt_no,omitempty"`
	// 匹配订单行编号
	SourcePoLineNo string `json:"source_po_line_no,omitempty" xml:"source_po_line_no,omitempty"`
	// 匹配订单头编号
	SourcePoNo string `json:"source_po_no,omitempty" xml:"source_po_no,omitempty"`
	// 应付发票行编号
	LineNo string `json:"line_no,omitempty" xml:"line_no,omitempty"`
}

var poolSupplierPreInvoiceItem = sync.Pool{
	New: func() any {
		return new(SupplierPreInvoiceItem)
	},
}

// GetSupplierPreInvoiceItem() 从对象池中获取SupplierPreInvoiceItem
func GetSupplierPreInvoiceItem() *SupplierPreInvoiceItem {
	return poolSupplierPreInvoiceItem.Get().(*SupplierPreInvoiceItem)
}

// ReleaseSupplierPreInvoiceItem 释放SupplierPreInvoiceItem
func ReleaseSupplierPreInvoiceItem(v *SupplierPreInvoiceItem) {
	v.LineCloseReason = ""
	v.LineStatus = ""
	v.SourceKpLineNo = ""
	v.SourceKpNo = ""
	v.LineRemark = ""
	v.TaxRate = ""
	v.TaxCode = ""
	v.TaxAmount = ""
	v.Amount = ""
	v.InvoiceLineQuantity = ""
	v.Spec = ""
	v.GoodsName = ""
	v.ItemCode = ""
	v.SourceInventory = ""
	v.SourceRtLineNo = ""
	v.SourceRtNo = ""
	v.SourcePoLineNo = ""
	v.SourcePoNo = ""
	v.LineNo = ""
	poolSupplierPreInvoiceItem.Put(v)
}
