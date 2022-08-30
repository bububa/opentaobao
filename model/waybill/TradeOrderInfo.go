package waybill

// TradeOrderInfo 结构体
type TradeOrderInfo struct {
	// 物流服务能力集合
	LogisticsServiceList []LogisticsService `json:"logistics_service_list,omitempty" xml:"logistics_service_list>logistics_service,omitempty"`
	// 包裹中的商品类型
	PackageItems []PackageItem `json:"package_items,omitempty" xml:"package_items>package_item,omitempty"`
	// 交易订单列表
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
	// 收货人
	ConsigneeName string `json:"consignee_name,omitempty" xml:"consignee_name,omitempty"`
	// 收货人联系方式
	ConsigneePhone string `json:"consignee_phone,omitempty" xml:"consignee_phone,omitempty"`
	// 订单渠道
	OrderChannelsType string `json:"order_channels_type,omitempty" xml:"order_channels_type,omitempty"`
	// 快递服务产品类型编码
	ProductType string `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// 发货人姓名
	SendName string `json:"send_name,omitempty" xml:"send_name,omitempty"`
	// 发货人联系方式
	SendPhone string `json:"send_phone,omitempty" xml:"send_phone,omitempty"`
	// 包裹号(或者ERP订单号)
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 收货人地址
	ConsigneeAddress *WaybillAddress `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
	// 使用者ID
	RealUserId int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
	// 包裹体积（立方厘米）
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 包裹重量（克）
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}
