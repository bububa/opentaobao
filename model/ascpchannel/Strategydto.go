package ascpchannel

import (
	"sync"
)

// Strategydto 结构体
type Strategydto struct {
	// 默认 4销售计划
	AicInventoryStrategyAvailableList []string `json:"aic_inventory_strategy_available_list,omitempty" xml:"aic_inventory_strategy_available_list>string,omitempty"`
	// 渠道出货规则
	ChannelPolicy string `json:"channel_policy,omitempty" xml:"channel_policy,omitempty"`
	// 渠道策略参数
	ChannelPolicyParam string `json:"channel_policy_param,omitempty" xml:"channel_policy_param,omitempty"`
	// 默认 4销售计划
	AicInventoryStrategy int64 `json:"aic_inventory_strategy,omitempty" xml:"aic_inventory_strategy,omitempty"`
}

var poolStrategydto = sync.Pool{
	New: func() any {
		return new(Strategydto)
	},
}

// GetStrategydto() 从对象池中获取Strategydto
func GetStrategydto() *Strategydto {
	return poolStrategydto.Get().(*Strategydto)
}

// ReleaseStrategydto 释放Strategydto
func ReleaseStrategydto(v *Strategydto) {
	v.AicInventoryStrategyAvailableList = v.AicInventoryStrategyAvailableList[:0]
	v.ChannelPolicy = ""
	v.ChannelPolicyParam = ""
	v.AicInventoryStrategy = 0
	poolStrategydto.Put(v)
}
