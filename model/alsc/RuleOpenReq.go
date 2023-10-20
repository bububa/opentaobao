package alsc

import (
	"sync"
)

// RuleOpenReq 结构体
type RuleOpenReq struct {
	// saas品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部id类型，wechat：微信openId  alipay：支付宝
	OuterType string `json:"outer_type,omitempty" xml:"outer_type,omitempty"`
	// 物理卡id
	PhysicalCardId string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
	// saas门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolRuleOpenReq = sync.Pool{
	New: func() any {
		return new(RuleOpenReq)
	},
}

// GetRuleOpenReq() 从对象池中获取RuleOpenReq
func GetRuleOpenReq() *RuleOpenReq {
	return poolRuleOpenReq.Get().(*RuleOpenReq)
}

// ReleaseRuleOpenReq 释放RuleOpenReq
func ReleaseRuleOpenReq(v *RuleOpenReq) {
	v.BrandId = ""
	v.Mobile = ""
	v.OuterId = ""
	v.OuterType = ""
	v.PhysicalCardId = ""
	v.ShopId = ""
	poolRuleOpenReq.Put(v)
}
