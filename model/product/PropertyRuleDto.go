package product

import (
	"sync"
)

// PropertyRuleDto 结构体
type PropertyRuleDto struct {
	// 字符串规则，正则表达式
	Pattern string `json:"pattern,omitempty" xml:"pattern,omitempty"`
	// 正则表达式,校验提示
	PatternMsg string `json:"pattern_msg,omitempty" xml:"pattern_msg,omitempty"`
	// 价格规则，最大值
	HighestPrice int64 `json:"highest_price,omitempty" xml:"highest_price,omitempty"`
	// 价格规则，最小值
	LowestPrice int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
	// 数字规则，最小值
	Min int64 `json:"min,omitempty" xml:"min,omitempty"`
	// 数字规则，最大值
	Max int64 `json:"max,omitempty" xml:"max,omitempty"`
	// 图片规则，最小长度
	MinLength int64 `json:"min_length,omitempty" xml:"min_length,omitempty"`
	// 图片规则，单图片最大体积
	MaxSize int64 `json:"max_size,omitempty" xml:"max_size,omitempty"`
	// 图片规则，最大图片数
	MaxCount int64 `json:"max_count,omitempty" xml:"max_count,omitempty"`
	// 字符串规则，最大长度
	MaxLength int64 `json:"max_length,omitempty" xml:"max_length,omitempty"`
	// 通用规则，是否必填
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
}

var poolPropertyRuleDto = sync.Pool{
	New: func() any {
		return new(PropertyRuleDto)
	},
}

// GetPropertyRuleDto() 从对象池中获取PropertyRuleDto
func GetPropertyRuleDto() *PropertyRuleDto {
	return poolPropertyRuleDto.Get().(*PropertyRuleDto)
}

// ReleasePropertyRuleDto 释放PropertyRuleDto
func ReleasePropertyRuleDto(v *PropertyRuleDto) {
	v.Pattern = ""
	v.PatternMsg = ""
	v.HighestPrice = 0
	v.LowestPrice = 0
	v.Min = 0
	v.Max = 0
	v.MinLength = 0
	v.MaxSize = 0
	v.MaxCount = 0
	v.MaxLength = 0
	v.Required = false
	poolPropertyRuleDto.Put(v)
}
