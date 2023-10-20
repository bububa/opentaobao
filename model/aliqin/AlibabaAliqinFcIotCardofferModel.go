package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotCardofferModel 结构体
type AlibabaAliqinFcIotCardofferModel struct {
	// 失效时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 生效时间
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// 订购时间
	OrderTime string `json:"order_time,omitempty" xml:"order_time,omitempty"`
	// 商品名称
	OfferName string `json:"offer_name,omitempty" xml:"offer_name,omitempty"`
	// 商品ID
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
}

var poolAlibabaAliqinFcIotCardofferModel = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotCardofferModel)
	},
}

// GetAlibabaAliqinFcIotCardofferModel() 从对象池中获取AlibabaAliqinFcIotCardofferModel
func GetAlibabaAliqinFcIotCardofferModel() *AlibabaAliqinFcIotCardofferModel {
	return poolAlibabaAliqinFcIotCardofferModel.Get().(*AlibabaAliqinFcIotCardofferModel)
}

// ReleaseAlibabaAliqinFcIotCardofferModel 释放AlibabaAliqinFcIotCardofferModel
func ReleaseAlibabaAliqinFcIotCardofferModel(v *AlibabaAliqinFcIotCardofferModel) {
	v.ExpireTime = ""
	v.EffectiveTime = ""
	v.OrderTime = ""
	v.OfferName = ""
	v.OfferId = ""
	poolAlibabaAliqinFcIotCardofferModel.Put(v)
}
