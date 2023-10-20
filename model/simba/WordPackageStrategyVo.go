package simba

import (
	"sync"
)

// WordPackageStrategyVo 结构体
type WordPackageStrategyVo struct {
	// 词包策略名称
	StrategyName string `json:"strategy_name,omitempty" xml:"strategy_name,omitempty"`
	// 关键词基础出价
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 词包策略,0:流量智选,1:好词优选,2:捡漏,2:卖点词包,3:类目优选
	StrategyId int64 `json:"strategy_id,omitempty" xml:"strategy_id,omitempty"`
	// 词包类型,0:流量智选,1:好词优选,19:捡漏,20:卖点词包,21:类目优选
	WordPackageType int64 `json:"word_package_type,omitempty" xml:"word_package_type,omitempty"`
	// 策略状态,1:在线,0:下线
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolWordPackageStrategyVo = sync.Pool{
	New: func() any {
		return new(WordPackageStrategyVo)
	},
}

// GetWordPackageStrategyVo() 从对象池中获取WordPackageStrategyVo
func GetWordPackageStrategyVo() *WordPackageStrategyVo {
	return poolWordPackageStrategyVo.Get().(*WordPackageStrategyVo)
}

// ReleaseWordPackageStrategyVo 释放WordPackageStrategyVo
func ReleaseWordPackageStrategyVo(v *WordPackageStrategyVo) {
	v.StrategyName = ""
	v.BidPrice = ""
	v.StrategyId = 0
	v.WordPackageType = 0
	v.OnlineStatus = 0
	v.AdgroupId = 0
	v.CampaignId = 0
	poolWordPackageStrategyVo.Put(v)
}
