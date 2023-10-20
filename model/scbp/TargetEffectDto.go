package scbp

import (
	"sync"
)

// TargetEffectDto 结构体
type TargetEffectDto struct {
	// 标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 选项值
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 日期(yyyy-MM-dd)
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 13:地域标签 14：人群标签
	TagRefType int64 `json:"tag_ref_type,omitempty" xml:"tag_ref_type,omitempty"`
	// 曝光
	Impr int64 `json:"impr,omitempty" xml:"impr,omitempty"`
	// 点击
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 消耗
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 推广时长
	OnlineMin int64 `json:"online_min,omitempty" xml:"online_min,omitempty"`
}

var poolTargetEffectDto = sync.Pool{
	New: func() any {
		return new(TargetEffectDto)
	},
}

// GetTargetEffectDto() 从对象池中获取TargetEffectDto
func GetTargetEffectDto() *TargetEffectDto {
	return poolTargetEffectDto.Get().(*TargetEffectDto)
}

// ReleaseTargetEffectDto 释放TargetEffectDto
func ReleaseTargetEffectDto(v *TargetEffectDto) {
	v.Subject = ""
	v.OptionValue = ""
	v.StatDate = ""
	v.TagRefType = 0
	v.Impr = 0
	v.Click = 0
	v.Cost = 0
	v.OnlineMin = 0
	poolTargetEffectDto.Put(v)
}
