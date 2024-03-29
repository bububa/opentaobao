package wdkitem

import (
	"sync"
)

// PromotionInfoDto 结构体
type PromotionInfoDto struct {
	// 展示文案
	DisplayText string `json:"display_text,omitempty" xml:"display_text,omitempty"`
	// 优惠类型, 0:减钱,1:打折,2:一口价,3:组合优惠
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 创建时间
	CreateDateTime string `json:"create_date_time,omitempty" xml:"create_date_time,omitempty"`
	// 活动名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 限购信息
	LimitInfo string `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 优惠金额，单位为分
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动类型, 1:单品活动,3:商品池活动
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
}

var poolPromotionInfoDto = sync.Pool{
	New: func() any {
		return new(PromotionInfoDto)
	},
}

// GetPromotionInfoDto() 从对象池中获取PromotionInfoDto
func GetPromotionInfoDto() *PromotionInfoDto {
	return poolPromotionInfoDto.Get().(*PromotionInfoDto)
}

// ReleasePromotionInfoDto 释放PromotionInfoDto
func ReleasePromotionInfoDto(v *PromotionInfoDto) {
	v.DisplayText = ""
	v.PromotionType = ""
	v.CreateDateTime = ""
	v.Name = ""
	v.StartTime = ""
	v.LimitInfo = ""
	v.EndTime = ""
	v.DiscountFee = 0
	v.ActivityId = 0
	v.ActivityType = 0
	poolPromotionInfoDto.Put(v)
}
