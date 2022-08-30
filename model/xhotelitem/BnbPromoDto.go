package xhotelitem

// BnbPromoDto 结构体
type BnbPromoDto struct {
	// 参与活动的rates
	RateInfos []Rateinfos `json:"rate_infos,omitempty" xml:"rate_infos>rateinfos,omitempty"`
	// 外部活动code
	OuterActivityCode string `json:"outer_activity_code,omitempty" xml:"outer_activity_code,omitempty"`
	// 优惠信息
	PromoInfo *PromoInfo `json:"promo_info,omitempty" xml:"promo_info,omitempty"`
}
