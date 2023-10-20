package tbk

import (
	"sync"
)

// MaifanPromotionDto 结构体
type MaifanPromotionDto struct {
	// 猫超买返卡活动结束时间
	MaifanPromotionEndTime string `json:"maifan_promotion_end_time,omitempty" xml:"maifan_promotion_end_time,omitempty"`
	// 猫超买返卡活动开始时间
	MaifanPromotionStartTime string `json:"maifan_promotion_start_time,omitempty" xml:"maifan_promotion_start_time,omitempty"`
	// 猫超买返卡面额
	MaifanPromotionDiscount string `json:"maifan_promotion_discount,omitempty" xml:"maifan_promotion_discount,omitempty"`
	// 猫超买返卡总数，-1代表不限量，其他大于等于0的值为总数
	MaifanPromotionCondition string `json:"maifan_promotion_condition,omitempty" xml:"maifan_promotion_condition,omitempty"`
}

var poolMaifanPromotionDto = sync.Pool{
	New: func() any {
		return new(MaifanPromotionDto)
	},
}

// GetMaifanPromotionDto() 从对象池中获取MaifanPromotionDto
func GetMaifanPromotionDto() *MaifanPromotionDto {
	return poolMaifanPromotionDto.Get().(*MaifanPromotionDto)
}

// ReleaseMaifanPromotionDto 释放MaifanPromotionDto
func ReleaseMaifanPromotionDto(v *MaifanPromotionDto) {
	v.MaifanPromotionEndTime = ""
	v.MaifanPromotionStartTime = ""
	v.MaifanPromotionDiscount = ""
	v.MaifanPromotionCondition = ""
	poolMaifanPromotionDto.Put(v)
}
