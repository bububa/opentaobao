package ascpchannel

// Strategydto 结构体
type Strategydto struct {
	// 默认 4
	AicInventoryStrategyAvailableList []int64 `json:"aic_inventory_strategy_available_list,omitempty" xml:"aic_inventory_strategy_available_list>int64,omitempty"`
	// 渠道出货规则
	ChannelPolicy string `json:"channel_policy,omitempty" xml:"channel_policy,omitempty"`
	// 渠道策略参数
	ChannelPolicyParam string `json:"channel_policy_param,omitempty" xml:"channel_policy_param,omitempty"`
	// 默认 4
	AicInventoryStrategy int64 `json:"aic_inventory_strategy,omitempty" xml:"aic_inventory_strategy,omitempty"`
}
