package ascpchannel

// BussinessInventoryLog 结构体
type BussinessInventoryLog struct {
	// 库存类型（包含在仓，不包含在途和物流差异
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 库存出入库数量
	ChangeQuantity string `json:"change_quantity,omitempty" xml:"change_quantity,omitempty"`
	// 货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// ERP单号
	ErpOrderCode string `json:"erp_order_code,omitempty" xml:"erp_order_code,omitempty"`
	// 单据类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 变动日期（格式YYYY-MM-DD HH:MI:SS）
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 业务主单号
	OpOrderId string `json:"op_order_id,omitempty" xml:"op_order_id,omitempty"`
	// 业务子单号
	OpSubOrderId string `json:"op_sub_order_id,omitempty" xml:"op_sub_order_id,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 物流LBX号
	WmsOrderCode string `json:"wms_order_code,omitempty" xml:"wms_order_code,omitempty"`
	// 淘系交易子单
	SubTradeId string `json:"sub_trade_id,omitempty" xml:"sub_trade_id,omitempty"`
	// 淘系交易主单
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 每日库存结余数量
	DaySurplusQuantity string `json:"day_surplus_quantity,omitempty" xml:"day_surplus_quantity,omitempty"`
}
