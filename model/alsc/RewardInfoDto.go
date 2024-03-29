package alsc

import (
	"sync"
)

// RewardInfoDto 结构体
type RewardInfoDto struct {
	// 扩展参数
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 奖励图标
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 奖励名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 奖品类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 百川抽奖结果
	UppPrizeResult string `json:"upp_prize_result,omitempty" xml:"upp_prize_result,omitempty"`
	// 奖品值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolRewardInfoDto = sync.Pool{
	New: func() any {
		return new(RewardInfoDto)
	},
}

// GetRewardInfoDto() 从对象池中获取RewardInfoDto
func GetRewardInfoDto() *RewardInfoDto {
	return poolRewardInfoDto.Get().(*RewardInfoDto)
}

// ReleaseRewardInfoDto 释放RewardInfoDto
func ReleaseRewardInfoDto(v *RewardInfoDto) {
	v.Ext = ""
	v.Icon = ""
	v.Name = ""
	v.RequestId = ""
	v.Type = ""
	v.UppPrizeResult = ""
	v.Value = ""
	poolRewardInfoDto.Put(v)
}
