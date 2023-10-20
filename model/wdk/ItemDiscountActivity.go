package wdk

import (
	"sync"
)

// ItemDiscountActivity 结构体
type ItemDiscountActivity struct {
	// 参加活动的渠道店ids
	ShopIds []string `json:"shop_ids,omitempty" xml:"shop_ids>string,omitempty"`
	// 优惠适用场景[APP|POS|POS+APP分别对应的值为1|2|1,2]
	Terminals []int64 `json:"terminals,omitempty" xml:"terminals>int64,omitempty"`
	// 商品特价优惠方式[itemDecreaseMoney:商品特价减钱;itemFixPrice:商品特价一口价;itemDiscount:商品特价打折]
	DiscountType string `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 活动名称,不超过10个英文字符
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 商家活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动详情描述,不超过30个英文字符
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 商家人群编码
	MerchantCrowdCode string `json:"merchant_crowd_code,omitempty" xml:"merchant_crowd_code,omitempty"`
	// 淘鲜达人群编码
	TxdCrowdCode string `json:"txd_crowd_code,omitempty" xml:"txd_crowd_code,omitempty"`
	// 渠道key
	ActivityChannel string `json:"activity_channel,omitempty" xml:"activity_channel,omitempty"`
	// 活动结束时间,时间戳
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动开始时间,时间戳
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 五道口活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 会员维度活动参与人群限制[-1:不限制;1:会员专享;2:非会员专享]
	MemberLimit int64 `json:"member_limit,omitempty" xml:"member_limit,omitempty"`
	// 周期优惠信息
	PeriodConfig *PeriodConfig `json:"period_config,omitempty" xml:"period_config,omitempty"`
	// 活动优先级，值越大表示优先级越高，必须大于0
	PriorityValue int64 `json:"priority_value,omitempty" xml:"priority_value,omitempty"`
	// coverBefore
	CoverBefore bool `json:"cover_before,omitempty" xml:"cover_before,omitempty"`
}

var poolItemDiscountActivity = sync.Pool{
	New: func() any {
		return new(ItemDiscountActivity)
	},
}

// GetItemDiscountActivity() 从对象池中获取ItemDiscountActivity
func GetItemDiscountActivity() *ItemDiscountActivity {
	return poolItemDiscountActivity.Get().(*ItemDiscountActivity)
}

// ReleaseItemDiscountActivity 释放ItemDiscountActivity
func ReleaseItemDiscountActivity(v *ItemDiscountActivity) {
	v.ShopIds = v.ShopIds[:0]
	v.Terminals = v.Terminals[:0]
	v.DiscountType = ""
	v.ActivityName = ""
	v.OutActId = ""
	v.Description = ""
	v.MerchantCrowdCode = ""
	v.TxdCrowdCode = ""
	v.ActivityChannel = ""
	v.EndTime = 0
	v.StartTime = 0
	v.ActivityId = 0
	v.MemberLimit = 0
	v.PeriodConfig = nil
	v.PriorityValue = 0
	v.CoverBefore = false
	poolItemDiscountActivity.Put(v)
}
