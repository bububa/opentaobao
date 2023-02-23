package ascp

// UpdateItemDistributionRequest 结构体
type UpdateItemDistributionRequest struct {
	// 指定分销商铺货详情
	DistributionInfoList []Long `json:"distribution_info_list,omitempty" xml:"distribution_info_list>long,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 【必传】要选择进行铺货的店铺宝贝 itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 要选择进行铺货的店铺宝贝 skuId， 如果没有sku就不传
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 币种
	DistributeCurrency string `json:"distribute_currency,omitempty" xml:"distribute_currency,omitempty"`
	// 币种
	RetailCurrency string `json:"retail_currency,omitempty" xml:"retail_currency,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 运费模板id， 可以通过alibaba.dchain.aoxiang.deliverytemplate.query 这个接口进行获取
	LogisticsCostTemplateId int64 `json:"logistics_cost_template_id,omitempty" xml:"logistics_cost_template_id,omitempty"`
	// 通用分销价格，卖给分销商的价格， 单位 分。 如果需要修改则需要传
	DistributePrice int64 `json:"distribute_price,omitempty" xml:"distribute_price,omitempty"`
	// 建议零售价。 建议分销商卖给消费者的价格， 单位 分。 非必传
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 1级分销价格
	Level1Price int64 `json:"level1_price,omitempty" xml:"level1_price,omitempty"`
	// 2级分销价格
	Level2Price int64 `json:"level2_price,omitempty" xml:"level2_price,omitempty"`
	// 3级分销价格
	Level3Price int64 `json:"level3_price,omitempty" xml:"level3_price,omitempty"`
	// 4级分销价格
	Level4Price int64 `json:"level4_price,omitempty" xml:"level4_price,omitempty"`
	// 5级分销价格
	Level5Price int64 `json:"level5_price,omitempty" xml:"level5_price,omitempty"`
}
