package alihouse

import (
	"sync"
)

// MarketingCouponDto 结构体
type MarketingCouponDto struct {
	// 使用楼盘列表
	ActivityProjectList []ProjectDetailInfoDto `json:"activity_project_list,omitempty" xml:"activity_project_list>project_detail_info_dto,omitempty"`
	// 活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 活动开始时间
	ActivityBeginTime string `json:"activity_begin_time,omitempty" xml:"activity_begin_time,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
	// 外部门店id，项目店和渠道标准店必传
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 券有效期
	ValidDate int64 `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 券金额
	CouponMoney int64 `json:"coupon_money,omitempty" xml:"coupon_money,omitempty"`
	// 券状态
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
	// 专车券返程金额
	CouponReturnMoney int64 `json:"coupon_return_money,omitempty" xml:"coupon_return_money,omitempty"`
	// 活动类型0-专车 2-往返
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
}

var poolMarketingCouponDto = sync.Pool{
	New: func() any {
		return new(MarketingCouponDto)
	},
}

// GetMarketingCouponDto() 从对象池中获取MarketingCouponDto
func GetMarketingCouponDto() *MarketingCouponDto {
	return poolMarketingCouponDto.Get().(*MarketingCouponDto)
}

// ReleaseMarketingCouponDto 释放MarketingCouponDto
func ReleaseMarketingCouponDto(v *MarketingCouponDto) {
	v.ActivityProjectList = v.ActivityProjectList[:0]
	v.ActivityEndTime = ""
	v.ActivityBeginTime = ""
	v.ActivityName = ""
	v.OuterActivityId = ""
	v.OuterStoreId = ""
	v.ValidDate = 0
	v.CouponMoney = 0
	v.ActivityStatus = 0
	v.CouponReturnMoney = 0
	v.ActivityType = 0
	poolMarketingCouponDto.Put(v)
}
