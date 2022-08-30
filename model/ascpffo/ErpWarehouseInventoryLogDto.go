package ascpffo

// ErpWarehouseInventoryLogDto 结构体
type ErpWarehouseInventoryLogDto struct {
	// 交易子单号
	BizSubTradeId string `json:"biz_sub_trade_id,omitempty" xml:"biz_sub_trade_id,omitempty"`
	// 交易主单号
	BizTradeId string `json:"biz_trade_id,omitempty" xml:"biz_trade_id,omitempty"`
	// 业务类型（PO0  普通采购,ADJ0 报废,ADJ1 盘点调整,ADJ2 状态调整,SO0  TOC销售,SO1  TOB销售,SO4  TOC补发,RTV0 普通采购退货,RSO0 TOC销售退货,SYS1 系统调账）
	BizActivityType string `json:"biz_activity_type,omitempty" xml:"biz_activity_type,omitempty"`
	// 业务子单号
	OperationDetailOrderId string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
	// 业务主单号
	OperationOrderId string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
	// 变化后占用数量
	ResultLockQuantity string `json:"result_lock_quantity,omitempty" xml:"result_lock_quantity,omitempty"`
	// 占用变化数量
	ChangeLockQuantity string `json:"change_lock_quantity,omitempty" xml:"change_lock_quantity,omitempty"`
	// 变化后库存数量
	ResultQuantity string `json:"result_quantity,omitempty" xml:"result_quantity,omitempty"`
	// 库存变化数量
	ChangeQuantity string `json:"change_quantity,omitempty" xml:"change_quantity,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 仓名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品条形码
	WhcBarCode string `json:"whc_bar_code,omitempty" xml:"whc_bar_code,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 仓内单据号
	WhOrderCode string `json:"wh_order_code,omitempty" xml:"wh_order_code,omitempty"`
	// 出入库时间
	OperateTime int64 `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 库存类型(1 良品，101 残品)
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
