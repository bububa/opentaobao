package scbp

import (
	"sync"
)

// TargetTagRecommendResultDto 结构体
type TargetTagRecommendResultDto struct {
	// 推荐标签名
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 推荐标签类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 高曝光竞价(单位：元)
	HighImprPrice string `json:"high_impr_price,omitempty" xml:"high_impr_price,omitempty"`
	// 推荐标签溢价值
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 推荐标签分数[0-100]
	Score int64 `json:"score,omitempty" xml:"score,omitempty"`
}

var poolTargetTagRecommendResultDto = sync.Pool{
	New: func() any {
		return new(TargetTagRecommendResultDto)
	},
}

// GetTargetTagRecommendResultDto() 从对象池中获取TargetTagRecommendResultDto
func GetTargetTagRecommendResultDto() *TargetTagRecommendResultDto {
	return poolTargetTagRecommendResultDto.Get().(*TargetTagRecommendResultDto)
}

// ReleaseTargetTagRecommendResultDto 释放TargetTagRecommendResultDto
func ReleaseTargetTagRecommendResultDto(v *TargetTagRecommendResultDto) {
	v.OptionValue = ""
	v.Type = ""
	v.HighImprPrice = ""
	v.Discount = 0
	v.Score = 0
	poolTargetTagRecommendResultDto.Put(v)
}
