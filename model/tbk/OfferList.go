package tbk

// OfferList 结构体
type OfferList struct {
	// 活动id
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 淘礼金领取URL，不支持使用短链接
	TljUrl string `json:"tlj_url,omitempty" xml:"tlj_url,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
