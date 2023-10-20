package pur

import (
	"sync"
)

// SupplierAsnItem 结构体
type SupplierAsnItem struct {
	// 关闭原因
	LineCloseReason string `json:"line_close_reason,omitempty" xml:"line_close_reason,omitempty"`
	// asn行状态
	LineStatus string `json:"line_status,omitempty" xml:"line_status,omitempty"`
	// asn行备注
	LineRemark string `json:"line_remark,omitempty" xml:"line_remark,omitempty"`
	// 木箱号
	Cass string `json:"cass,omitempty" xml:"cass,omitempty"`
	// 集装箱号
	ContainerNo string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	// 装箱单号
	PackingListNo string `json:"packing_list_no,omitempty" xml:"packing_list_no,omitempty"`
	// 发货金额
	DeliveryAmount string `json:"delivery_amount,omitempty" xml:"delivery_amount,omitempty"`
	// 发货数量
	DeliveryQty string `json:"delivery_qty,omitempty" xml:"delivery_qty,omitempty"`
	// 购买方式
	ProcurementMethod string `json:"procurement_method,omitempty" xml:"procurement_method,omitempty"`
	// 发货批次号
	ShippingBatchNo string `json:"shipping_batch_no,omitempty" xml:"shipping_batch_no,omitempty"`
	// asn来源订单行号
	SourcePoLineNum string `json:"source_po_line_num,omitempty" xml:"source_po_line_num,omitempty"`
	// asn来源订单号
	SourcePoNo string `json:"source_po_no,omitempty" xml:"source_po_no,omitempty"`
	// 根据订单发货标识
	IsPoSource string `json:"is_po_source,omitempty" xml:"is_po_source,omitempty"`
	// asn单行编号
	AsnLineNum string `json:"asn_line_num,omitempty" xml:"asn_line_num,omitempty"`
	// 延期原因
	DelayReason string `json:"delay_reason,omitempty" xml:"delay_reason,omitempty"`
}

var poolSupplierAsnItem = sync.Pool{
	New: func() any {
		return new(SupplierAsnItem)
	},
}

// GetSupplierAsnItem() 从对象池中获取SupplierAsnItem
func GetSupplierAsnItem() *SupplierAsnItem {
	return poolSupplierAsnItem.Get().(*SupplierAsnItem)
}

// ReleaseSupplierAsnItem 释放SupplierAsnItem
func ReleaseSupplierAsnItem(v *SupplierAsnItem) {
	v.LineCloseReason = ""
	v.LineStatus = ""
	v.LineRemark = ""
	v.Cass = ""
	v.ContainerNo = ""
	v.PackingListNo = ""
	v.DeliveryAmount = ""
	v.DeliveryQty = ""
	v.ProcurementMethod = ""
	v.ShippingBatchNo = ""
	v.SourcePoLineNum = ""
	v.SourcePoNo = ""
	v.IsPoSource = ""
	v.AsnLineNum = ""
	v.DelayReason = ""
	poolSupplierAsnItem.Put(v)
}
