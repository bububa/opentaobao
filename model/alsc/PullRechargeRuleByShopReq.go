package alsc

import (
	"sync"
)

// PullRechargeRuleByShopReq 结构体
type PullRechargeRuleByShopReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 会员卡，礼品卡模板id，选填，不填则默认查询会员卡
	CardTemplateId string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	// 上次更新时间戳
	MaxUpdateTime string `json:"max_update_time,omitempty" xml:"max_update_time,omitempty"`
	// 外部品牌id,brand_id和out_brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店id,shop_id和out_shop_id不可同时为空
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolPullRechargeRuleByShopReq = sync.Pool{
	New: func() any {
		return new(PullRechargeRuleByShopReq)
	},
}

// GetPullRechargeRuleByShopReq() 从对象池中获取PullRechargeRuleByShopReq
func GetPullRechargeRuleByShopReq() *PullRechargeRuleByShopReq {
	return poolPullRechargeRuleByShopReq.Get().(*PullRechargeRuleByShopReq)
}

// ReleasePullRechargeRuleByShopReq 释放PullRechargeRuleByShopReq
func ReleasePullRechargeRuleByShopReq(v *PullRechargeRuleByShopReq) {
	v.BrandId = ""
	v.CardTemplateId = ""
	v.MaxUpdateTime = ""
	v.OutBrandId = ""
	v.OutShopId = ""
	v.ShopId = ""
	poolPullRechargeRuleByShopReq.Put(v)
}
