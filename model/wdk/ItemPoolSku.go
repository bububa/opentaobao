package wdk

// ItemPoolSku 结构体
type ItemPoolSku struct {
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 换购价格
	ExchangePrice int64 `json:"exchange_price,omitempty" xml:"exchange_price,omitempty"`
	// 换购限量
	ExchangeTotalLimit int64 `json:"exchange_total_limit,omitempty" xml:"exchange_total_limit,omitempty"`
	// 逻辑分组ID
	LogicGroupNumber int64 `json:"logic_group_number,omitempty" xml:"logic_group_number,omitempty"`
	// 是否为一口价
	IsFixPrice bool `json:"is_fix_price,omitempty" xml:"is_fix_price,omitempty"`
	// 一口价金额【分】
	FixPrice int64 `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
	// 是否为商品折扣
	IsDiscountRate bool `json:"is_discount_rate,omitempty" xml:"is_discount_rate,omitempty"`
	// 折扣率，300=3折
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 是否为减钱
	IsDecreaseMoney bool `json:"is_decrease_money,omitempty" xml:"is_decrease_money,omitempty"`
	// 减钱【元】
	DecreaseMoney int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
	// 淘宝item和shop的对应关系， k-itemId, v-shopId
	ItemShopRelation string `json:"item_shop_relation,omitempty" xml:"item_shop_relation,omitempty"`
}
