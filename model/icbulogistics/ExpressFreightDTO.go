package icbulogistics

// ExpressFreightDto 结构体
type ExpressFreightDto struct {
	// 费用项列表
	ExpressQuoteItemList []ExpressQuoteItemList `json:"express_quote_item_list,omitempty" xml:"express_quote_item_list>express_quote_item_list,omitempty"`
	// 销售总价
	SalesAmount string `json:"sales_amount,omitempty" xml:"sales_amount,omitempty"`
	// 折扣总价
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 仓库信息
	Warehouse *WarehouseDto `json:"warehouse,omitempty" xml:"warehouse,omitempty"`
	// （废弃为空！！面单通过alibaba.onetouch.logistics.express.order.detail.get获取）原条码PDF Base64编码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 物流订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 上门揽收信息
	PickupInfo *PickupInfoDto `json:"pickup_info,omitempty" xml:"pickup_info,omitempty"`
}
