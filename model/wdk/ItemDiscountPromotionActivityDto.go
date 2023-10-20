package wdk

import (
	"sync"
)

// ItemDiscountPromotionActivityDto 结构体
type ItemDiscountPromotionActivityDto struct {
	// 活动周几生效
	Weekdays []string `json:"weekdays,omitempty" xml:"weekdays>string,omitempty"`
	// 活动每天生效时间段
	EveryDayPeriods []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
	// 优惠适用场景:1:APP  2:POS
	Terminals []string `json:"terminals,omitempty" xml:"terminals>string,omitempty"`
	// 门店列表
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 外部门店列表
	OuterStoreIds []string `json:"outer_store_ids,omitempty" xml:"outer_store_ids>string,omitempty"`
	// 人群编码,saas平台人群编码:NEW_USER: 新用户, OLD_USER：老用户，LIGHT_NEW_USER：闪购新客
	MemberCrowdCodes []string `json:"member_crowd_codes,omitempty" xml:"member_crowd_codes>string,omitempty"`
	// 外部订单号
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 操作人ID
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人姓名
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 单品活动类型1:减钱 2:一口价 3:打折
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 营销活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 活动开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 限购信息
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
}

var poolItemDiscountPromotionActivityDto = sync.Pool{
	New: func() any {
		return new(ItemDiscountPromotionActivityDto)
	},
}

// GetItemDiscountPromotionActivityDto() 从对象池中获取ItemDiscountPromotionActivityDto
func GetItemDiscountPromotionActivityDto() *ItemDiscountPromotionActivityDto {
	return poolItemDiscountPromotionActivityDto.Get().(*ItemDiscountPromotionActivityDto)
}

// ReleaseItemDiscountPromotionActivityDto 释放ItemDiscountPromotionActivityDto
func ReleaseItemDiscountPromotionActivityDto(v *ItemDiscountPromotionActivityDto) {
	v.Weekdays = v.Weekdays[:0]
	v.EveryDayPeriods = v.EveryDayPeriods[:0]
	v.Terminals = v.Terminals[:0]
	v.StoreIds = v.StoreIds[:0]
	v.OuterStoreIds = v.OuterStoreIds[:0]
	v.MemberCrowdCodes = v.MemberCrowdCodes[:0]
	v.OutActId = ""
	v.ActivityName = ""
	v.Description = ""
	v.CreatorId = ""
	v.CreatorName = ""
	v.DiscountType = 0
	v.ActId = 0
	v.StartTime = 0
	v.EndTime = 0
	v.Limit = nil
	poolItemDiscountPromotionActivityDto.Put(v)
}
