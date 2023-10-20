package ascp

import (
	"sync"
)

// OrderLine 结构体
type OrderLine struct {
	// 批次列表
	Batchs []AlibabaDchainAoxiangWmsDeliveryorderConfirmBatch `json:"batchs,omitempty" xml:"batchs>alibaba_dchain_aoxiang_wms_deliveryorder_confirm_batch,omitempty"`
	// 交易主单号
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 交易子单号
	SubSourceOrderCode string `json:"sub_source_order_code,omitempty" xml:"sub_source_order_code,omitempty"`
	// 货主编码。ERP对接场景下，传ERP内部的商家编码或者淘系的sellerId
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 单据行号
	OrderLineNo string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
	// 平台交易订单编码
	OrderSourceCode string `json:"order_source_code,omitempty" xml:"order_source_code,omitempty"`
	// 平台交易子订单编码
	SubSourceCode string `json:"sub_source_code,omitempty" xml:"sub_source_code,omitempty"`
	// 商品仓储系统编码
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 交易平台商品编码
	ExtCode string `json:"ext_code,omitempty" xml:"ext_code,omitempty"`
	// 实发商品数量
	ActualQty string `json:"actual_qty,omitempty" xml:"actual_qty,omitempty"`
	// 批次编号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 生产日期(YYYY-MM-DD)
	ProductDate string `json:"product_date,omitempty" xml:"product_date,omitempty"`
	// 过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 生产批号
	ProduceCode string `json:"produce_code,omitempty" xml:"produce_code,omitempty"`
	// 商品的二维码(类似电子产品的SN码;用来进行商品的溯源;多个二维码之间用分号;隔开)
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 货品sn编码
	SnCode string `json:"sn_code,omitempty" xml:"sn_code,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 供应商编码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 支付平台交易号(淘系订单传支付宝交易号)
	PayNo string `json:"pay_no,omitempty" xml:"pay_no,omitempty"`
	// 零售价(零售价=实际成交价+单件商品折扣金额)
	RetailPrice string `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 实际成交价
	ActualPrice string `json:"actual_price,omitempty" xml:"actual_price,omitempty"`
	// 单件商品折扣金额
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 货品编码id，必填
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 应发货品数量
	PlanQty int64 `json:"plan_qty,omitempty" xml:"plan_qty,omitempty"`
	// sn列表
	SnList *SnList `json:"sn_list,omitempty" xml:"sn_list,omitempty"`
}

var poolOrderLine = sync.Pool{
	New: func() any {
		return new(OrderLine)
	},
}

// GetOrderLine() 从对象池中获取OrderLine
func GetOrderLine() *OrderLine {
	return poolOrderLine.Get().(*OrderLine)
}

// ReleaseOrderLine 释放OrderLine
func ReleaseOrderLine(v *OrderLine) {
	v.Batchs = v.Batchs[:0]
	v.SourceOrderCode = ""
	v.SubSourceOrderCode = ""
	v.OwnerCode = ""
	v.ItemCode = ""
	v.OrderLineNo = ""
	v.OrderSourceCode = ""
	v.SubSourceCode = ""
	v.ItemId = ""
	v.InventoryType = ""
	v.ItemName = ""
	v.ExtCode = ""
	v.ActualQty = ""
	v.BatchCode = ""
	v.ProductDate = ""
	v.ExpireDate = ""
	v.ProduceCode = ""
	v.QrCode = ""
	v.SnCode = ""
	v.SupplierName = ""
	v.SupplierCode = ""
	v.PayNo = ""
	v.RetailPrice = ""
	v.ActualPrice = ""
	v.DiscountAmount = ""
	v.ScItemId = ""
	v.PlanQty = 0
	v.SnList = nil
	poolOrderLine.Put(v)
}
