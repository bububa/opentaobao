package bus

import (
	"sync"
)

// TopStandardInsRecommendRequest 结构体
type TopStandardInsRecommendRequest struct {
	// OFF_LINE（线下自助机）、OTA（OTA）、SUBMIT_ORDER_PAGE（下单页）、RECOMMEND_PAGE （推荐弹层）、STORE_WINDOW（商家窗口）。目前仅开放了线下自助机和商家窗口。
	TargetChannel string `json:"target_channel,omitempty" xml:"target_channel,omitempty"`
	// 站点信息对象
	StationInfo *TopInsStationInfo `json:"station_info,omitempty" xml:"station_info,omitempty"`
	// 商户身份对象
	MerchantInfo *TopInsMerchantInfo `json:"merchant_info,omitempty" xml:"merchant_info,omitempty"`
	// 商品信息对象
	CommodityInfo *TopInsCommodityInfo `json:"commodity_info,omitempty" xml:"commodity_info,omitempty"`
}

var poolTopStandardInsRecommendRequest = sync.Pool{
	New: func() any {
		return new(TopStandardInsRecommendRequest)
	},
}

// GetTopStandardInsRecommendRequest() 从对象池中获取TopStandardInsRecommendRequest
func GetTopStandardInsRecommendRequest() *TopStandardInsRecommendRequest {
	return poolTopStandardInsRecommendRequest.Get().(*TopStandardInsRecommendRequest)
}

// ReleaseTopStandardInsRecommendRequest 释放TopStandardInsRecommendRequest
func ReleaseTopStandardInsRecommendRequest(v *TopStandardInsRecommendRequest) {
	v.TargetChannel = ""
	v.StationInfo = nil
	v.MerchantInfo = nil
	v.CommodityInfo = nil
	poolTopStandardInsRecommendRequest.Put(v)
}
