package promotion

import (
	"sync"
)

// PosKeyPromotionRuleDetailDto 结构体
type PosKeyPromotionRuleDetailDto struct {
	// 扩展属性
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 1,9,12,15四个等级
	UnitPoint string `json:"unit_point,omitempty" xml:"unit_point,omitempty"`
	// 数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 价格
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

var poolPosKeyPromotionRuleDetailDto = sync.Pool{
	New: func() any {
		return new(PosKeyPromotionRuleDetailDto)
	},
}

// GetPosKeyPromotionRuleDetailDto() 从对象池中获取PosKeyPromotionRuleDetailDto
func GetPosKeyPromotionRuleDetailDto() *PosKeyPromotionRuleDetailDto {
	return poolPosKeyPromotionRuleDetailDto.Get().(*PosKeyPromotionRuleDetailDto)
}

// ReleasePosKeyPromotionRuleDetailDto 释放PosKeyPromotionRuleDetailDto
func ReleasePosKeyPromotionRuleDetailDto(v *PosKeyPromotionRuleDetailDto) {
	v.ExtMap = ""
	v.UnitPoint = ""
	v.Num = 0
	v.UnitPrice = 0
	poolPosKeyPromotionRuleDetailDto.Put(v)
}
