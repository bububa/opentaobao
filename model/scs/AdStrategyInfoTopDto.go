package scs

// AdStrategyInfoTopDto 结构体
type AdStrategyInfoTopDto struct {
	// 商品列表，小于等于10个
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 拉新快场景中的新客定义，支持&#34;1,365&#34;,&#34;1,500&#34;等
	BehaviorWindowsList []string `json:"behavior_windows_list,omitempty" xml:"behavior_windows_list>string,omitempty"`
	// 子场景
	MarketScene string `json:"market_scene,omitempty" xml:"market_scene,omitempty"`
	// 优化目标，1036,促进购买;1037,促进进店;1038,促进收藏加购；入会快新会员场景只能使用1041 促进入会量；入会快老会员场景只能使用1036 1038；
	MarketAim string `json:"market_aim,omitempty" xml:"market_aim,omitempty"`
	// 套餐包模式计划需要填的预算，最小值2000
	OrderAmount string `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// 套餐包模式计划需要填的预算，最小值2000
	Budget string `json:"budget,omitempty" xml:"budget,omitempty"`
	// 投放模式，效率优先effect_first，消耗优先cost_first
	LaunchStrategyType string `json:"launch_strategy_type,omitempty" xml:"launch_strategy_type,omitempty"`
	// 预估周期
	PredictCycle int64 `json:"predict_cycle,omitempty" xml:"predict_cycle,omitempty"`
	// 侧重人群开关，默认为1
	CrowdMode int64 `json:"crowd_mode,omitempty" xml:"crowd_mode,omitempty"`
}
