package ascpffo

// WarehouseInventoryLogQueryDto 结构体
type WarehouseInventoryLogQueryDto struct {
	// 业务类型(PO0  普通采购,ADJ0 报废,ADJ1 盘点调整,ADJ2 状态调整,SO0  TOC销售,SO1  TOB销售,SO4  TOC补发,RTV0 普通采购退货,RSO0 TOC销售退货,SYS1 系统调账)
	BizActivityType string `json:"biz_activity_type,omitempty" xml:"biz_activity_type,omitempty"`
	// 交易主单号
	BizTradeId string `json:"biz_trade_id,omitempty" xml:"biz_trade_id,omitempty"`
	// 业务主单号
	OperationOrderId string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 出入库开始时间，ms级
	GmtCreateEnd int64 `json:"gmt_create_end,omitempty" xml:"gmt_create_end,omitempty"`
	// 出入库截止时间，ms级
	GmtCreateStart int64 `json:"gmt_create_start,omitempty" xml:"gmt_create_start,omitempty"`
	// 库存类型(1 良品，101 残品)
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小，最大200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
