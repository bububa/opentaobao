package bus

// TopStandardInsRecommendRequest 结构体
type TopStandardInsRecommendRequest struct {
	// 站点信息对象
	StationInfo *TopInsStationInfo `json:"station_info,omitempty" xml:"station_info,omitempty"`
	// OFF_LINE（线下自助机）、OTA（OTA）、SUBMIT_ORDER_PAGE（下单页）、RECOMMEND_PAGE （推荐弹层）、STORE_WINDOW（商家窗口）。目前仅开放了线下自助机和商家窗口。
	TargetChannel string `json:"target_channel,omitempty" xml:"target_channel,omitempty"`
	// 商户身份对象
	MerchantInfo *TopInsMerchantInfo `json:"merchant_info,omitempty" xml:"merchant_info,omitempty"`
	// 商品信息对象
	CommodityInfo *TopInsCommodityInfo `json:"commodity_info,omitempty" xml:"commodity_info,omitempty"`
}
