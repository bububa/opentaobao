package aedropshiper

// AeopOrderInfo 结构体
type AeopOrderInfo struct {
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 物流状态
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 订单金额
	OrderAmount *SimpleMoney `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// 子订单列表
	ChildOrderList []AeopChildOrderInfo `json:"child_order_list,omitempty" xml:"child_order_list>aeop_child_order_info,omitempty"`
	// 订单物流信息列表
	LogisticsInfoList []AeopOrderLogisticsInfo `json:"logistics_info_list,omitempty" xml:"logistics_info_list>aeop_order_logistics_info,omitempty"`
	// 店铺信息
	StoreInfo *AeopStoreInfo `json:"store_info,omitempty" xml:"store_info,omitempty"`
}
