package lstmarketing

import (
	"sync"
)

// Promotiondtolist 结构体
type Promotiondtolist struct {
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 营销类型
	PromotionTypeName string `json:"promotion_type_name,omitempty" xml:"promotion_type_name,omitempty"`
	// 优惠金额，分为单位
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
}

var poolPromotiondtolist = sync.Pool{
	New: func() any {
		return new(Promotiondtolist)
	},
}

// GetPromotiondtolist() 从对象池中获取Promotiondtolist
func GetPromotiondtolist() *Promotiondtolist {
	return poolPromotiondtolist.Get().(*Promotiondtolist)
}

// ReleasePromotiondtolist 释放Promotiondtolist
func ReleasePromotiondtolist(v *Promotiondtolist) {
	v.ActivityId = ""
	v.ActivityName = ""
	v.PromotionTypeName = ""
	v.DiscountFee = 0
	poolPromotiondtolist.Put(v)
}
