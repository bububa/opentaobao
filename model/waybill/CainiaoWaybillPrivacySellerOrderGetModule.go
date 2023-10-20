package waybill

import (
	"sync"
)

// CainiaoWaybillPrivacySellerOrderGetModule 结构体
type CainiaoWaybillPrivacySellerOrderGetModule struct {
	// 商家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 日期
	OrderDate string `json:"order_date,omitempty" xml:"order_date,omitempty"`
	// 订单渠道
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 店铺id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 隐私次数
	PrivacyCount int64 `json:"privacy_count,omitempty" xml:"privacy_count,omitempty"`
}

var poolCainiaoWaybillPrivacySellerOrderGetModule = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillPrivacySellerOrderGetModule)
	},
}

// GetCainiaoWaybillPrivacySellerOrderGetModule() 从对象池中获取CainiaoWaybillPrivacySellerOrderGetModule
func GetCainiaoWaybillPrivacySellerOrderGetModule() *CainiaoWaybillPrivacySellerOrderGetModule {
	return poolCainiaoWaybillPrivacySellerOrderGetModule.Get().(*CainiaoWaybillPrivacySellerOrderGetModule)
}

// ReleaseCainiaoWaybillPrivacySellerOrderGetModule 释放CainiaoWaybillPrivacySellerOrderGetModule
func ReleaseCainiaoWaybillPrivacySellerOrderGetModule(v *CainiaoWaybillPrivacySellerOrderGetModule) {
	v.SellerId = ""
	v.OrderDate = ""
	v.OrderChannel = ""
	v.ShopId = ""
	v.PrivacyCount = 0
	poolCainiaoWaybillPrivacySellerOrderGetModule.Put(v)
}
