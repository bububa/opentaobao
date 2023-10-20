package qimen

import (
	"sync"
)

// InventoryRule 结构体
type InventoryRule struct {
	// channelRatioRules
	ChannelRatioRules []ChannelRatioRule `json:"channelRatioRules,omitempty" xml:"channelRatioRules>channel_ratio_rule,omitempty"`
	// 奇门仓储字段,C1223,string(50),,
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 奇门仓储字段,C1223,string(50),,
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
}

var poolInventoryRule = sync.Pool{
	New: func() any {
		return new(InventoryRule)
	},
}

// GetInventoryRule() 从对象池中获取InventoryRule
func GetInventoryRule() *InventoryRule {
	return poolInventoryRule.Get().(*InventoryRule)
}

// ReleaseInventoryRule 释放InventoryRule
func ReleaseInventoryRule(v *InventoryRule) {
	v.ChannelRatioRules = v.ChannelRatioRules[:0]
	v.ActionType = ""
	v.ItemCode = ""
	poolInventoryRule.Put(v)
}
