package ascp

// CreateItemDistributionRequest 结构体
type CreateItemDistributionRequest struct {
	// 【必传】要选择进行铺货的店铺宝贝 itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 【非必传】如果商品没有sku则必传，如果有sku则非必传。
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 【必传】商品名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 要选择进行铺货的店铺宝贝 skuId， 如果没有sku就不传
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku的商家编码， 如果没有sku就不传， 如果有sku则必传
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// sku规格
	SkuTitle string `json:"sku_title,omitempty" xml:"sku_title,omitempty"`
	// 【必传】货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 币种
	DistributeCurrency string `json:"distribute_currency,omitempty" xml:"distribute_currency,omitempty"`
	// 币种
	RetailCurrency string `json:"retail_currency,omitempty" xml:"retail_currency,omitempty"`
	// 【必传】运费模板id， 可以通过alibaba.dchain.aoxiang.deliverytemplate.query 这个接口进行获取
	LogisticsCostTemplateId int64 `json:"logistics_cost_template_id,omitempty" xml:"logistics_cost_template_id,omitempty"`
	// 【必传】分销价格，卖给分销商的价格， 单位 分
	DistributePrice int64 `json:"distribute_price,omitempty" xml:"distribute_price,omitempty"`
	// 建议零售价。 建议分销商卖给消费者的价格， 单位 分。 非必传
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
