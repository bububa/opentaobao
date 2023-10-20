package xhotelonlineorder

import (
	"sync"
)

// TopHotelPromotion 结构体
type TopHotelPromotion struct {
	// 订单命中的优惠计算规则描述(由优惠计算规则解析后的文本形式)
	RuleDesc string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// 活动优惠描述
	PromotionDesc string `json:"promotion_desc,omitempty" xml:"promotion_desc,omitempty"`
	// 活动Icon文案描述
	ActivityIconDesc string `json:"activity_icon_desc,omitempty" xml:"activity_icon_desc,omitempty"`
	// 活动分组名称
	ActivityGroupName string `json:"activity_group_name,omitempty" xml:"activity_group_name,omitempty"`
	// 对应的航旅平台活动规则
	ActivityCode string `json:"activity_code,omitempty" xml:"activity_code,omitempty"`
	// 使用的活动规则code
	RuleCode string `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	// 卖家系统中的活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 消費里程数
	ConsumeMileage int64 `json:"consume_mileage,omitempty" xml:"consume_mileage,omitempty"`
	// 活动使用的优惠金额，以分为单位
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 活动使用的卖家优惠金额，以分为单位
	SellerAmount int64 `json:"seller_amount,omitempty" xml:"seller_amount,omitempty"`
	// 活动使用优惠次数
	Times int64 `json:"times,omitempty" xml:"times,omitempty"`
	// 活动展示优先级（越大优先级越高）
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 优惠的呈现方式,目前有如下定义:,1. 金额类优惠,2. 使用次数类优惠,3. 权益类优惠,4. 后返类优惠
	AmountType int64 `json:"amount_type,omitempty" xml:"amount_type,omitempty"`
	// 活动使用的折扣
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 优惠计算规则类型
	RuleType int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 优惠维度
	RuleTarget int64 `json:"rule_target,omitempty" xml:"rule_target,omitempty"`
	// 出资方是否卖家还是平台
	InvestorType int64 `json:"investor_type,omitempty" xml:"investor_type,omitempty"`
	// 是否&#34;积分加钱购&#34;类型活动
	IntegralActivity bool `json:"integral_activity,omitempty" xml:"integral_activity,omitempty"`
}

var poolTopHotelPromotion = sync.Pool{
	New: func() any {
		return new(TopHotelPromotion)
	},
}

// GetTopHotelPromotion() 从对象池中获取TopHotelPromotion
func GetTopHotelPromotion() *TopHotelPromotion {
	return poolTopHotelPromotion.Get().(*TopHotelPromotion)
}

// ReleaseTopHotelPromotion 释放TopHotelPromotion
func ReleaseTopHotelPromotion(v *TopHotelPromotion) {
	v.RuleDesc = ""
	v.PromotionDesc = ""
	v.ActivityIconDesc = ""
	v.ActivityGroupName = ""
	v.ActivityCode = ""
	v.RuleCode = ""
	v.ActivityName = ""
	v.ConsumeMileage = 0
	v.Amount = 0
	v.SellerAmount = 0
	v.Times = 0
	v.Priority = 0
	v.Tid = 0
	v.AmountType = 0
	v.Discount = 0
	v.RuleType = 0
	v.RuleTarget = 0
	v.InvestorType = 0
	v.IntegralActivity = false
	poolTopHotelPromotion.Put(v)
}
