package aliqin

// AlibabaaliqinfciotcardofferModel 结构体
type AlibabaaliqinfciotcardofferModel struct {
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
