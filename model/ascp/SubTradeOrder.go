package ascp

// SubTradeOrder 结构体
type SubTradeOrder struct {
	// 交易子单ID
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 商品宝贝ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品的最小库存单位Sku的id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 组合货品比例
	CombineScItemRatio string `json:"combine_sc_item_ratio,omitempty" xml:"combine_sc_item_ratio,omitempty"`
	// 组合货品code
	CombineScItemCode string `json:"combine_sc_item_code,omitempty" xml:"combine_sc_item_code,omitempty"`
	// 子订单类型:标示该子交易单来源交易，还是翱象增加的，枚举值(00=交易，10=翱象绑定)
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 主赠品标识，主品=0，赠品=1
	ItemConsignType string `json:"item_consign_type,omitempty" xml:"item_consign_type,omitempty"`
	// 活动ID，赠品绑赠的活动，赠品回滚需要字段
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 货品仓储code
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 币种，USD-美元，CNY-人民币，RUB-卢布，JPY-日元，EUR-欧元，GBP-英镑，HKD-港币，NZD-新西兰元，SGD-新加坡元，AUD-澳元，KRW-韩元，THB-泰铢
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 发货类型CN=菜鸟发货，SC的商家仓发货
	ConsignType string `json:"consign_type,omitempty" xml:"consign_type,omitempty"`
	// 赠品ID，赠品回滚需要字段
	GiftId string `json:"gift_id,omitempty" xml:"gift_id,omitempty"`
	// 交易子单的来源主交易单号。如果是赠品行则表示赠品的来源主交易单号
	ParentTradeId string `json:"parent_trade_id,omitempty" xml:"parent_trade_id,omitempty"`
	// 赠品子单号。如果有值的话与tradeId相同
	GiftOrderId string `json:"gift_order_id,omitempty" xml:"gift_order_id,omitempty"`
	// 仓库货品编码
	WarehouseScItemCode string `json:"warehouse_sc_item_code,omitempty" xml:"warehouse_sc_item_code,omitempty"`
	// 货品条码
	ScItemBarCode string `json:"sc_item_bar_code,omitempty" xml:"sc_item_bar_code,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 货主ID
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 枚举值： 1 = 组合货品 2 = 非组合货品
	CombineType string `json:"combine_type,omitempty" xml:"combine_type,omitempty"`
	// 应发数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 子单金额；单位:分，如:20007，表示:20元7分
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
}
