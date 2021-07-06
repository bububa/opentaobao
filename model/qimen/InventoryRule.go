package qimen

// InventoryRule 结构体
type InventoryRule struct {
	// channelRatioRules
	ChannelRatioRules []ChannelRatioRule `json:"channelRatioRules,omitempty" xml:"channelRatioRules>channel_ratio_rule,omitempty"`
	// 奇门仓储字段,C1223,string(50),,
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 奇门仓储字段,C1223,string(50),,
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
}
