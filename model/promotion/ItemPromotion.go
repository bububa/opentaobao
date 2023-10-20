package promotion

import (
	"sync"
)

// ItemPromotion 结构体
type ItemPromotion struct {
	// 活动名称。
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动的有效条件、人群和行为描述。
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 活动开始时间。
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间。
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 用户标签。当is_user_tag为true时，该值才有意义。
	UserTag string `json:"user_tag,omitempty" xml:"user_tag,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动范围：0表示全部参与； 1表示部分商品参与。
	ParticipateRange int64 `json:"participate_range,omitempty" xml:"participate_range,omitempty"`
	// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
	DecreaseAmount int64 `json:"decrease_amount,omitempty" xml:"decrease_amount,omitempty"`
	// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 是否指定用户标签。
	IsUserTag bool `json:"is_user_tag,omitempty" xml:"is_user_tag,omitempty"`
	// 是否有减钱行为。
	IsDecreaseMoney bool `json:"is_decrease_money,omitempty" xml:"is_decrease_money,omitempty"`
	// 是否有打折行为。
	IsDiscount bool `json:"is_discount,omitempty" xml:"is_discount,omitempty"`
}

var poolItemPromotion = sync.Pool{
	New: func() any {
		return new(ItemPromotion)
	},
}

// GetItemPromotion() 从对象池中获取ItemPromotion
func GetItemPromotion() *ItemPromotion {
	return poolItemPromotion.Get().(*ItemPromotion)
}

// ReleaseItemPromotion 释放ItemPromotion
func ReleaseItemPromotion(v *ItemPromotion) {
	v.Name = ""
	v.Description = ""
	v.StartTime = ""
	v.EndTime = ""
	v.UserTag = ""
	v.ActivityId = 0
	v.ParticipateRange = 0
	v.DecreaseAmount = 0
	v.DiscountRate = 0
	v.IsUserTag = false
	v.IsDecreaseMoney = false
	v.IsDiscount = false
	poolItemPromotion.Put(v)
}
