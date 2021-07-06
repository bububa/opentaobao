package nlife

// Traderefundlist 结构体
type Traderefundlist struct {
	// 采购退货单号
	TradeRefundNo string `json:"trade_refund_no,omitempty" xml:"trade_refund_no,omitempty"`
	// 采购退货的供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 采购退货单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 采购退货的供应商id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 门店Id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
