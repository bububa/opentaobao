package alsc

import (
	"sync"
)

// CustomerGetOpenReq 结构体
type CustomerGetOpenReq struct {
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
	// 顾客id
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
}

var poolCustomerGetOpenReq = sync.Pool{
	New: func() any {
		return new(CustomerGetOpenReq)
	},
}

// GetCustomerGetOpenReq() 从对象池中获取CustomerGetOpenReq
func GetCustomerGetOpenReq() *CustomerGetOpenReq {
	return poolCustomerGetOpenReq.Get().(*CustomerGetOpenReq)
}

// ReleaseCustomerGetOpenReq 释放CustomerGetOpenReq
func ReleaseCustomerGetOpenReq(v *CustomerGetOpenReq) {
	v.BrandId = ""
	v.Mobile = ""
	v.OuterId = ""
	v.OuterType = ""
	v.PhysicalCardId = ""
	v.ShopId = ""
	v.CustomerId = ""
	poolCustomerGetOpenReq.Put(v)
}
