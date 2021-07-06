package wms

// CainiaoReturnBillReturnorderinfo 结构体
type CainiaoReturnBillReturnorderinfo struct {
	// 订单商品信息列表
	OrderItemList []CainiaoReturnBillOrderitemlist `json:"order_item_list,omitempty" xml:"order_item_list>cainiao_return_bill_orderitemlist,omitempty"`
	// ERP订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓库订单编码，LBX号
	CnOrderCode string `json:"cn_order_code,omitempty" xml:"cn_order_code,omitempty"`
	// 仓库订单完成时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 此销退订单对应原销售订单的菜鸟订单号
	PreCnOrderCode string `json:"pre_cn_order_code,omitempty" xml:"pre_cn_order_code,omitempty"`
	// 单据类型： 501 退货入库
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
}
