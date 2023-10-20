package scbp

import (
	"sync"
)

// AccountEffectDto 结构体
type AccountEffectDto struct {
	// 日期(yyyy-MM-dd)
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 曝光
	Impr string `json:"impr,omitempty" xml:"impr,omitempty"`
	// 点击
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 消耗
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 推广时长
	OnlineMin string `json:"online_min,omitempty" xml:"online_min,omitempty"`
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 平均点击花费
	ClickCostAvg string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 曝光
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 单位小时，保留一位小数，例如13.5表示13.5小时
	OnlineTime string `json:"online_time,omitempty" xml:"online_time,omitempty"`
}

var poolAccountEffectDto = sync.Pool{
	New: func() any {
		return new(AccountEffectDto)
	},
}

// GetAccountEffectDto() 从对象池中获取AccountEffectDto
func GetAccountEffectDto() *AccountEffectDto {
	return poolAccountEffectDto.Get().(*AccountEffectDto)
}

// ReleaseAccountEffectDto 释放AccountEffectDto
func ReleaseAccountEffectDto(v *AccountEffectDto) {
	v.StatDate = ""
	v.Impr = ""
	v.Click = ""
	v.Cost = ""
	v.OnlineMin = ""
	v.ClickCnt = ""
	v.ClickCostAvg = ""
	v.Ctr = ""
	v.ImpressionCnt = ""
	v.OnlineTime = ""
	poolAccountEffectDto.Put(v)
}
