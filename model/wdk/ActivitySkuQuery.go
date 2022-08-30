package wdk

// ActivitySkuQuery 结构体
type ActivitySkuQuery struct {
	// 需要查询的商品skuCodes
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
	// 自定义同步的渠道配置
	ChannelConfigList []ChannelConfig `json:"channel_config_list,omitempty" xml:"channel_config_list>channel_config,omitempty"`
	// 商家活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 五道口活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 分页参数
	Page *BasePageQuery `json:"page,omitempty" xml:"page,omitempty"`
	// 是否自定义渠道同步
	ByChannel bool `json:"by_channel,omitempty" xml:"by_channel,omitempty"`
}
