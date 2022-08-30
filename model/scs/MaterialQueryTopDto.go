package scs

// MaterialQueryTopDto 结构体
type MaterialQueryTopDto struct {
	// 场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 各个场景的对应值为：各个场景对应的值为：拉新快--拉新快204，首单直降205，派样拉新216，入会快--入会拉新213，老会员激活213，货品加速--货品加速211，上新快--行业新品207，新品首降214，预热蓄水--预热蓄水219
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 页码
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
}
