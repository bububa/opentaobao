package wdk

// ItemPoolSkuActivityElementDto 结构体
type ItemPoolSkuActivityElementDto struct {
	// 商品渠道配置信息
	SkuChannelConfigs []SkuChannelConfigDto `json:"sku_channel_configs,omitempty" xml:"sku_channel_configs>sku_channel_config_dto,omitempty"`
	// 商品编码，与商品条码必选其一，或者同时传入以商品条码为准
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品条码，与商品编码必选其一，或者同时传入以商品条码为准
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 操作人ID
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人姓名
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 换购价（单位分）
	ExchangePrice int64 `json:"exchange_price,omitempty" xml:"exchange_price,omitempty"`
	// 营销活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 商品限购
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
}
