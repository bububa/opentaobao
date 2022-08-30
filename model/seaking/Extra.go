package seaking

// Extra 结构体
type Extra struct {
	// 商品所在平台
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 商品id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 店铺id
	SubIdentifier string `json:"sub_identifier,omitempty" xml:"sub_identifier,omitempty"`
	// 店铺所在平台
	SubIdentifyType string `json:"sub_identify_type,omitempty" xml:"sub_identify_type,omitempty"`
}
