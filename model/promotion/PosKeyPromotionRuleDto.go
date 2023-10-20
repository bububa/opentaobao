package promotion

import (
	"sync"
)

// PosKeyPromotionRuleDto 结构体
type PosKeyPromotionRuleDto struct {
	// 优惠规则
	DetailList []PosKeyPromotionRuleDetailDto `json:"detail_list,omitempty" xml:"detail_list>pos_key_promotion_rule_detail_dto,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// postkey
	PosKey int64 `json:"pos_key,omitempty" xml:"pos_key,omitempty"`
}

var poolPosKeyPromotionRuleDto = sync.Pool{
	New: func() any {
		return new(PosKeyPromotionRuleDto)
	},
}

// GetPosKeyPromotionRuleDto() 从对象池中获取PosKeyPromotionRuleDto
func GetPosKeyPromotionRuleDto() *PosKeyPromotionRuleDto {
	return poolPosKeyPromotionRuleDto.Get().(*PosKeyPromotionRuleDto)
}

// ReleasePosKeyPromotionRuleDto 释放PosKeyPromotionRuleDto
func ReleasePosKeyPromotionRuleDto(v *PosKeyPromotionRuleDto) {
	v.DetailList = v.DetailList[:0]
	v.Name = ""
	v.PosKey = 0
	poolPosKeyPromotionRuleDto.Put(v)
}
