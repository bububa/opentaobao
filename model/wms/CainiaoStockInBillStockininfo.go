package wms

// CainiaoStockInBillStockininfo 结构体
type CainiaoStockInBillStockininfo struct {
	// ERP订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 菜鸟订单编码
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// 单据类型：  904 普通入库 306 B2B入库单 601 采购入库
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 仓库订单完成时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 入库单状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败WMS_CONFIRMED 仓库生产完成
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 入库单信息
	OrderItemList []CainiaoStockInBillOrderitemlist `json:"order_item_list,omitempty" xml:"order_item_list>cainiao_stock_in_bill_orderitemlist,omitempty"`
}
