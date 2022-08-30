package ascp

// ConsignOrder 结构体
type ConsignOrder struct {
	// 交易子单列表
	SubTradeOrders []SubTradeOrder `json:"sub_trade_orders,omitempty" xml:"sub_trade_orders>sub_trade_order,omitempty"`
	// 淘系交易主单号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 翱象的发货单号
	ConsignOrderCode string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty"`
	// 仓储单号，推送到仓内的单号，自动流转订单才有
	WmsOrderCode string `json:"wms_order_code,omitempty" xml:"wms_order_code,omitempty"`
	// 物流单状态，0=待审核/1=仓接单/2=取消
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 店铺，交易成单的店铺ID
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 货主，发货的货主ID信息
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 发货仓
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 发货配
	DeliveryCps string `json:"delivery_cps,omitempty" xml:"delivery_cps,omitempty"`
	// 物流服务；里面多个值时用英文逗号隔开 201为驿站送货上门服务 202为顺丰配送服务
	AsdpAds string `json:"asdp_ads,omitempty" xml:"asdp_ads,omitempty"`
	// 买家留言
	BuyerMessage string `json:"buyer_message,omitempty" xml:"buyer_message,omitempty"`
	// 卖家留言
	SellerMessage string `json:"seller_message,omitempty" xml:"seller_message,omitempty"`
	// 0=无拆合，1=拆单，2=合单，3=拆合单
	AssemblyType string `json:"assembly_type,omitempty" xml:"assembly_type,omitempty"`
	// 合单包含的交易单号，多交易单号逗号隔开
	MergeTradeIds string `json:"merge_trade_ids,omitempty" xml:"merge_trade_ids,omitempty"`
	// 异常场景：value=-1；鲲鹏店铺的订单ERP需要进行Hold单； value=0；ERP无需做任何特殊处理 value=1；交易状态：禁止发货；ERP需要进行Hold单； value=2；交易状态：待卖家发货；ERP需要进行Hold单； value=3；交易状态：待卖家发货；ERP开发进行发货处理；
	ErpHold string `json:"erp_hold,omitempty" xml:"erp_hold,omitempty"`
	// 是否自动流转，0=否，1=是
	AutoFlow int64 `json:"auto_flow,omitempty" xml:"auto_flow,omitempty"`
	// 要求的发货时间
	PlanDeliveryTime int64 `json:"plan_delivery_time,omitempty" xml:"plan_delivery_time,omitempty"`
	// 要求的送达时间
	PlanSignTime int64 `json:"plan_sign_time,omitempty" xml:"plan_sign_time,omitempty"`
	// 发货单创建时间
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 发货单修改时间
	UpdateTime int64 `json:"update_time,omitempty" xml:"update_time,omitempty"`
}
