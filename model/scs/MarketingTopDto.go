package scs

// MarketingTopDto 结构体
type MarketingTopDto struct {
	// 子场景，拉新快adStrategyDkx：1023拉新快，1079首单直降，1113派样拉新；入会快adStrategyRuHui：1095入会拉新，1096老会员激活；adStrategyShangXin：1023行业新品，1098新品首降；adStrategyProductSpeed：1203货品加速；活动预热 adStrategyYuRe ： 1114 活动预热
	MarketScene int64 `json:"market_scene,omitempty" xml:"market_scene,omitempty"`
	// 优化目标，1036,促进购买;1037,促进进店;1038,促进收藏加购；入会快新会员场景只能使用1041 促进入会量；入会快老会员场景只能使用1036 1038；
	MarketAim int64 `json:"market_aim,omitempty" xml:"market_aim,omitempty"`
}
